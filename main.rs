use std::{thread, time};
use rand::seq::SliceRandom;
use reqwest::blocking::Client;
use core::time::Duration;

struct Config {
    domains: Vec<&'static str>,
    interval_secs: u64,
    timeout_secs: u64,
}

fn ping_domains(config: Config) {
    let client = Client::builder()
        .timeout(Duration::from_secs(config.timeout_secs))
        .build()
        .unwrap();

    loop {
        let domain = config.domains.choose(&mut rand::thread_rng()).unwrap();

        match client.get(*domain).send() {
            Ok(res) => println!("Pinged {}: {:?}", domain, res.status()),
            Err(e) => println!("Failed to ping {}: {}", domain, e),
        }

        thread::sleep(time::Duration::from_secs(config.interval_secs));
    }
}


fn main() {
   ping_domains(Config {
        domains: vec!["https://www.google.com", "https://www.bing.com"],
        interval_secs: 1,
        timeout_secs: 5,
    })
}
