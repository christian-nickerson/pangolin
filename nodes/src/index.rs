mod config;

use config::settings::Settings;

fn main() {
    // load config
    let _settings = match Settings::load() {
        Ok(settings) => {
            println!("configuration loaded");
            settings
        }
        Err(e) => {
            eprintln!("config load failed: {}", e);
            eprintln!("reverting to default config");
            Settings::default()
        }
    };
}
