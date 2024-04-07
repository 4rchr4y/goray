use axum::async_trait;
use chrono::Utc;
use std::sync::Arc;
use surrealdb::{engine::remote::ws::Client, Surreal};

use crate::domain::{
    self,
    model::project::{NewProjectInput, Project},
};

#[derive(Debug)]
pub struct ProjectRepositoryImpl {
    client: Arc<Surreal<Client>>,
    table_name: String,
}

impl ProjectRepositoryImpl {
    pub fn new(client: Arc<Surreal<Client>>, table: &str) -> Self {
        Self {
            client,
            table_name: table.into(),
        }
    }
}

#[async_trait]
impl domain::port::ProjectRepository for ProjectRepositoryImpl {
    async fn create_project(&self, input: NewProjectInput) -> Result<Vec<Project>, domain::Error> {
        let result: Vec<Project> = self
            .client
            .create(&self.table_name)
            .content(Project {
                id: None,
                path: input.path,
                updated: Utc::now().timestamp(),
            })
            .await?;

        Ok(result)
    }

    async fn delete_by_id(&self, id: String) -> Result<Option<Project>, domain::Error> {
        let result: Option<Project> = self.client.delete((&self.table_name, id)).await?;

        Ok(result)
    }
}