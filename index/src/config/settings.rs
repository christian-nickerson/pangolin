use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Settings {
    pub server: Server,
    pub transformers: Transformers,
    pub spacy: Spacy,
    pub metadata: Metadata,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Server {
    pub api: ApiServer,
    pub embeddings: EmbeddingsServer,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ApiServer {
    pub name: String,
    pub port: u16,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct EmbeddingsServer {
    pub name: String,
    pub port: u16,
    pub shutdown_period: u32,
    pub worker_threads: u32,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Transformers {
    pub model_list: Vec<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Spacy {
    pub model_list: Vec<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Metadata {
    pub database: Database,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Database {
    pub r#type: String,
    pub host: String,
    pub port: u16,
    pub dbname: String,
    pub username: String,
    pub password: String,
}

impl Default for Settings {
    fn default() -> Self {
        Self {
            server: Server {
                api: ApiServer {
                    name: "Pangolin".to_string(),
                    port: 3000,
                },
                embeddings: EmbeddingsServer {
                    name: "EmbeddingService".to_string(),
                    port: 50051,
                    shutdown_period: 5,
                    worker_threads: 10,
                },
            },
            transformers: Transformers {
                model_list: vec![
                    "all-mpnet-base-v2".to_string(),
                    "all-MiniLM-L6-v2".to_string(),
                ],
            },
            spacy: Spacy {
                model_list: vec![],
            },
            metadata: Metadata {
                database: Database {
                    r#type: "sqlite".to_string(),
                    host: "localhost".to_string(),
                    port: 5432,
                    dbname: "test".to_string(),
                    username: "postgres".to_string(),
                    password: "postgres".to_string(),
                },
            },
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_default_settings() {
        let settings = Settings::default();
        assert_eq!(settings.server.api.name, "Pangolin");
        assert_eq!(settings.server.api.port, 3000);
        assert_eq!(settings.server.embeddings.port, 50051);
        assert_eq!(settings.metadata.database.r#type, "sqlite");
    }

    #[test]
    fn test_settings_serialization() {
        let settings = Settings::default();
        let serialized = toml::to_string(&settings);
        assert!(serialized.is_ok());
    }
}
