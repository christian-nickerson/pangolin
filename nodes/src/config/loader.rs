use config::{Config, ConfigError, Environment, File};
use std::path::Path;

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
}

#[cfg(test)]
mod tests {
    use std::path::Path;

    #[test]
    fn test_load_settings_from_file() {
        // Test loading from the actual settings.toml file if it exists
        if Path::new("../../settings.toml").exists() {
            let settings = crate::config::settings::Settings::from_file("../../settings.toml");
            assert!(settings.is_ok());
        }
    }
}
