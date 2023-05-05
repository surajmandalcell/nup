use core::time::Duration;
use rand::seq::SliceRandom;
use reqwest::blocking::Client;

use std::{thread, time};

use crate::types::config::Config;

pub fn ping(config: Config) {
    let client = Client::builder()
        .timeout(Duration::from_secs(config.timeout_secs))
        .build()
        .unwrap();

    loop {
        let domain = config.domains.choose(&mut rand::thread_rng()).unwrap();
        let start_time = time::Instant::now();

        match client.get(*domain).send() {
            Ok(res) => {
                let mut log: String = format!("{}", domain);
                if config.flag_status {
                    log = format!("{} | Status: {}", log, res.status());
                }
                if config.flag_latency {
                    let elapsed = start_time.elapsed();
                    log = format!(
                        "{} | Time: {}.{:03}s",
                        log,
                        elapsed.as_secs(),
                        elapsed.subsec_millis().to_string()
                    );
                }
                if config.flag_verbose {
                    println!("{}", log)
                };
            }
            Err(_err) => {
                if config.flag_verbose {
                    println!("{} | Status: Failed", domain)
                }
            }
        }

        thread::sleep(time::Duration::from_secs(config.interval_secs));
    }
}
