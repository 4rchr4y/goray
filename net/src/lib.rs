mod domain;
mod infra;

pub mod config;

use std::sync::Arc;

use domain::model::error::Error;
use tokio_util::sync::CancellationToken as TokioCancellationToken;
use tower::ServiceBuilder;
use tower_http::{
    compression::{
        predicate::{NotForContentType, SizeAbove},
        CompressionLayer, Predicate,
    },
    request_id::MakeRequestUuid,
    ServiceBuilderExt,
};

use crate::{
    config::CONF,
    domain::{model::OptionExtension, service::ServiceRoot},
};

#[macro_use]
extern crate async_trait;

#[macro_use]
extern crate serde;
extern crate serde_json;

#[macro_use]
extern crate tracing;

#[macro_use]
extern crate mac;

const MIX_COMPRESS_SIZE: u16 = 512; // TODO: this value should be used from a net_conf.toml file

pub async fn bind(_: TokioCancellationToken) -> Result<(), Error> {
    let conf = CONF
        .get()
        .ok_or_config_invalid("Configuration was not defined", None)?;

    let service = ServiceBuilder::new()
        .catch_panic()
        .set_x_request_id(MakeRequestUuid)
        .propagate_x_request_id();

    let service_hub = ServiceRoot::new(conf);
    let schema = infra::graphql::build_schema(service_hub);
    let service = service.layer(
        CompressionLayer::new().compress_when(
            SizeAbove::new(MIX_COMPRESS_SIZE) // don't compress below 512 bytes
                .and(NotForContentType::IMAGES), // don't compress images
        ),
    );

    let router = infra::web::router(service, schema); // TODO: consider to use Cow<T>

    info!("Moss has been successfully launched on {}", conf.bind);

    axum_server::bind(conf.bind)
        .serve(router.clone().into_make_service())
        .await?;

    Ok(())
}
