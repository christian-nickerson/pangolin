use config::{Config, ConfigError, Environment, File};
use std::path::Path;

impl crate::config::settings::Settings {
    /// Load settings from a TOML file
    pub fn from_file<P: AsRef<Path>>(path: P) -> Result<Self, ConfigError> {
        let config = Config::builder()
            .add_source(File::from(path.as_ref()))
            .add_source(
                Environment::with_prefix("PANGOLIN")
                    .prefix_separator("_")
                    .separator("__"),
            )
            .build()?;

        config.try_deserialize()
    }

    /// Load settings from the default settings.toml file
    pub fn load() -> Result<Self, ConfigError> {
        Self::from_file("settings.toml")
    }
}

#[cfg(test)]
mod tests {
    use std::env;
    use std::path::Path;

    #[test]
    fn test_load_settings_from_file() {
        if Path::new("../../settings.toml").exists() {
            let settings = crate::config::settings::Settings::from_file("settings.toml");
            assert!(settings.is_ok());

            assert_eq!(settings.unwrap().server.api.port, 3000);
        }
    }

    #[test]
    #[serial_test::serial]
    fn test_env_var_overrides_toml_settings() {
        unsafe {
            env::set_var("PANGOLIN_server__API__PORT", "9999");
            env::set_var("PANGOLIN_SERVER__EMBEDDINGS__NAME", "OverriddenEmbedding");
        }

        let settings = crate::config::settings::Settings::from_file("settings.toml")
            .expect("failed to read settings file");

        assert_eq!(settings.server.api.port, 9999);
        assert_eq!(settings.server.embeddings.name, "OverriddenEmbedding");

        // Clean up
        unsafe {
            env::remove_var("PANGOLIN_SERVER__API__PORT");
            env::remove_var("PANGOLIN_SERVER__EMBEDDINGS__NAME");
        }
    }
}
