use std::sync::Arc;

use crate::domain::{
    self,
    model::project::{NewProjectInput, Project},
    port::ProjectRepository,
};

#[derive(Debug, Clone)]
pub struct ProjectService {
    pub repo: Arc<dyn ProjectRepository>,
}

impl ProjectService {
    pub fn new(repo: Arc<dyn ProjectRepository>) -> Self {
        Self { repo }
    }

    pub async fn create_project(
        &self,
        input: NewProjectInput,
    ) -> Result<Vec<Project>, domain::Error> {
        let result = self.repo.create_project(input).await?;

        Ok(result)
    }

    pub async fn delete_by_id(&self, id: String) -> Result<Project, domain::Error> {
        let result = self.repo.delete_by_id(id).await?;

        Ok(result.unwrap())
    }
}
