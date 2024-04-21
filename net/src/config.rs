use conf::pref::Preference;
use sea_orm::DatabaseConnection;
use std::{
    net::SocketAddr,
    sync::{Arc, OnceLock},
};

pub static CONF: OnceLock<Config> = OnceLock::new();

#[derive(Clone, Debug)]
pub struct Config {
    pub bind: SocketAddr,
    pub preference: Arc<Preference>,
    pub conn: Arc<DatabaseConnection>,
}