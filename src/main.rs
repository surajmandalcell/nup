mod lstatic;
mod types;

use types::Config;

use core::time::Duration;
use rand::seq::SliceRandom;
use reqwest::blocking::Client;

use std::{thread, time};

fn ping(config: Config) {
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
                println!("{}", log);
            }
            Err(_err) => println!("Failed to ping {}: {}", domain, _err),
        }

        thread::sleep(time::Duration::from_secs(config.interval_secs));
    }
}

fn main() {
    let mut arg_latency = false;
    let mut arg_status = false;

    for arg in std::env::args().skip(1) {
        match arg.as_str() {
            "-t" => arg_latency = true,
            "-s" => arg_status = true,
            "-h" | "--help" => lstatic::statics::help_msg(),
            _ => (),
        }
    }

    ping(Config {
        domains: lstatic::constants::DOMAINS.to_vec(),
        interval_secs: 1,
        timeout_secs: 5,
        flag_status: arg_status,
        flag_latency: arg_latency,
    })
}
