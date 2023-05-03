pub struct Args {
    pub latency: bool,
    pub status: bool,
}

pub struct Config {
    pub domains: Vec<&'static str>,
    pub interval_secs: u64,
    pub timeout_secs: u64,
    pub flag_latency: bool,
    pub flag_status: bool,
}
