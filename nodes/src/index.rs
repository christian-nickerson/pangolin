mod config;

use config::settings::Settings;
use log::{info, warn};

fn main() {
    // load config
    let _settings = match Settings::load() {
        Ok(settings) => {
            info!("configuration loaded");
            settings
        }
        Err(e) => {
            warn!("config load failed: {}", e);
            warn!("reverting to default config");
            Settings::default()
        }
    };
}
