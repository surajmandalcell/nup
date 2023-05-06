package app

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"nup/types"
)

func Ping(config types.Config) {
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
