pub mod constants {
    pub const DOMAINS: [&str; 2] = ["https://www.google.com", "https://www.bing.com"];
}

pub mod statics {
    use std::process::exit;

    pub fn help_msg() {
        let help_msg = indoc::indoc! {"
        \nUsage: ping [OPTION]

        Options:
            -t          Show latency
            -s          Show status code
            -h, --help  Show this help message
    "};
        println!("{}", help_msg);
        exit(0);
    }
}
