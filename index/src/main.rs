mod config;

use config::Settings;

fn main() {
    // Load configuration from settings.toml
    let settings = match Settings::load() {
        Ok(settings) => {
            println!("Successfully loaded configuration:");
            println!("  API Server: {} on port {}", settings.server.api.name, settings.server.api.port);
            println!("  Embeddings Server: {} on port {}", settings.server.embeddings.name, settings.server.embeddings.port);
            println!("  Database: {} at {}:{}", settings.metadata.database.r#type, settings.metadata.database.host, settings.metadata.database.port);
            println!("  Transformer models: {:?}", settings.transformers.model_list);
            settings
        }
        Err(e) => {
            eprintln!("Failed to load configuration: {}", e);
            eprintln!("Using default configuration...");
            Settings::default()
        }
    };

    println!("Pangolin Index application started with configuration loaded!");
}
