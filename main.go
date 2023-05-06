package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Domains      []string
	IntervalSecs uint64
	TimeoutSecs  uint64
	FlagLatency  bool
	FlagStatus   bool
	FlagVerbose  bool
}

type Args struct {
	Latency bool
	Status  bool
	Verbose bool
}

func main() {
	args := Args{
		Latency: false,
		Status:  false,
		Verbose: false,
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-s":
			args.Status = true
		case "-t":
			args.Latency = true
		case "-v":
			args.Verbose = true
		case "-h", "--help":
			helpMsg()
			return
		}
	}

	ping(Config{
		Domains:      []string{"https://www.google.com", "https://www.bing.com"},
		IntervalSecs: 1,
		TimeoutSecs:  5,
		FlagStatus:   args.Status,
		FlagLatency:  args.Latency,
		FlagVerbose:  args.Verbose,
	})
}

func ping(config Config) {
	client := &http.Client{
		Timeout: time.Duration(config.TimeoutSecs) * time.Second,
	}

	for {
		domain := config.Domains[rand.Intn(len(config.Domains))]
		startTime := time.Now()

		resp, err := client.Get(domain)
		if err != nil {
			if config.FlagVerbose {
				fmt.Printf("%s | Status: Failed\n", domain)
			}
			continue
		}

		log := fmt.Sprintf("%s", domain)
		if config.FlagStatus {
			log = fmt.Sprintf("%s | Status: %s", log, resp.Status)
		}
		if config.FlagLatency {
			elapsed := time.Since(startTime)
			log = fmt.Sprintf("%s | Time: %ds.%03ds", log, int64(elapsed.Seconds()), elapsed.Milliseconds())
		}
		if config.FlagVerbose {
			fmt.Println(log)
		}

		time.Sleep(time.Duration(config.IntervalSecs) * time.Second)
	}
}

func helpMsg() {
	fmt.Printf(`
	Usage: nup [OPTION]

	Options:
		-t          Show latency
		-s          Show status code
		-h, --help  Show this help message
	`)
}
