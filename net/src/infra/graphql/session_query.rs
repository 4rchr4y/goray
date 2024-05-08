use async_graphql::{Context, Object, Result as GraphqlResult};
use chrono::{Duration, Utc};
use common::id::NanoId;
use graphql_utl::GraphQLExtendError;
use http::HeaderMap;
use tokio::sync::RwLock;

use crate::domain::{
    model::{
        error::Error,
        session::{CreateSessionInput, Session, SessionInfo},
    },
    service::{project_service::ProjectService, session_service::SessionService},
};

#[derive(Default)]
pub(super) struct SessionMutation;

#[Object]
impl SessionMutation {
    async fn create_session(
        &self,
        ctx: &Context<'_>,
        input: CreateSessionInput,
    ) -> GraphqlResult<SessionInfo> {
        let session_service = ctx.data::<RwLock<SessionService>>()?;
        let project_service = ctx.data::<RwLock<Option<ProjectService>>>()?;

        let mut session_service_lock = session_service.write().await;
        let session = session_service_lock
            .create_session(&input, project_service)
            .await
            .extend_error()?;

        Ok(session)
    }

    async fn restore_session(
        &self,
        ctx: &Context<'_>,
        session_id: NanoId,
    ) -> GraphqlResult<Session> {
        let project_service = ctx.data::<RwLock<Option<ProjectService>>>()?;
        let mut session_service_lock = ctx.data::<RwLock<SessionService>>()?.write().await;
        let session = session_service_lock
            .restore_session(session_id, project_service)
            .await
            .extend_error()?;

        Ok(session)
    }

    #[graphql(name = "getRecentSessions")]
    #[graphql_mac::require_header("session-id")]
    async fn get_recent(
        &self,
        ctx: &Context<'_>,
        #[graphql(default_with = "(Utc::now() - Duration::days(30)).timestamp()")] start_time: i64,
        #[graphql(validator(minimum = 1, maximum = 10), default = 10)] limit: u64,
    ) -> GraphqlResult<Vec<Session>> {
        let session_service_lock = ctx.data::<RwLock<SessionService>>()?.write().await;
        let result = session_service_lock
            .get_recent_list(start_time, limit)
            .await
            .extend_error()?;

        Ok(result)
    }
}
