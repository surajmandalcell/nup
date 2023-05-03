pub mod constants {
    use std::env::var;

    pub const DOMAINS: [&str; 2] = ["https://www.google.com", "https://www.bing.com"];

    #[allow(dead_code)]
    pub fn get_config_path() -> String {
        const CONFIG_PATH_TMP: &str = "/.config/nup/config.json";
        let home: String = var("HOME").unwrap_or_else(|_| "/tmp".to_string());
        let config_path: String = format!("{}{}", CONFIG_PATH_TMP, home);
        return config_path;
    }
}

pub mod statics {
    use std::process::exit;

    pub fn help_msg() {
        let help_msg = indoc::indoc! {"
        \nUsage: nup [OPTION]

        Options:
            -t          Show latency
            -s          Show status code
            -h, --help  Show this help message
    "};
        println!("{}", help_msg);
        exit(0);
    }
}
