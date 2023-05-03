mod ping;
mod statics;
mod types;

use crate::ping::ping;
use types::args::Args;
use types::config::Config;

fn main() {
    let mut args = Args {
        latency: false,
        status: false,
        // To be implemented
        // verbose: false,
    };

    for arg in std::env::args().skip(1) {
        match arg.as_str() {
            "-s" => args.status = true,
            "-t" => args.latency = true,
            "-h" | "--help" => statics::statics::help_msg(),
            _ => (),
        }
    }

    ping(Config {
        domains: statics::constants::DOMAINS.to_vec(),
        interval_secs: 1,
        timeout_secs: 5,
        flag_status: args.status,
        flag_latency: args.latency,
    })
}
