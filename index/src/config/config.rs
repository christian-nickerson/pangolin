use config::{Config, ConfigError, Environment, File};
use std::path::Path;
use std::env;


impl crate::config::settings::Settings {
    /// Load settings from a TOML file
    pub fn from_file<P: AsRef<Path>>(path: P) -> Result<Self, ConfigError> {
        let config = Config::builder()
            .add_source(File::from(path.as_ref()))
            .add_source(Environment::with_prefix("PANGOLIN"))
            .build()?;

        config.try_deserialize()
    }

    /// Load settings from the default settings.toml file
    pub fn load() -> Result<Self, ConfigError> {
        Self::from_file("settings.toml")
    }

    /// Load settings with custom configuration
    pub fn load_with_config(config: Config) -> Result<Self, ConfigError> {
        config.try_deserialize()
    }
}



#[cfg(test)]
mod tests {
    use std::path::Path;
    use std::env;
    use tempfile::NamedTempFile;
    use std::io::Write;

    #[test]
    fn test_env_override() {
        let mut file = NamedTempFile::new().unwrap();
        writeln!(file, "[server.api]\nname = \"test-server\"\nport = 1234").unwrap();

        env::set_var("PANGOLIN_SERVER__API__PORT", "5678");
        let settings = crate::config::settings::Settings::from_file(file.path()).unwrap();
        assert_eq!(settings.server.api.port, 5678);
        env::remove_var("PANGOLIN_SERVER__API__PORT");
    }

    #[test]
    fn test_load_settings_from_file() {
        // Test loading from the actual settings.toml file if it exists
        if Path::new("../../settings.toml").exists() {
            let settings = crate::config::settings::Settings::from_file("../../settings.toml");
            assert!(settings.is_ok());
        }
    }
}
