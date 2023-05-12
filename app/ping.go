package app

import (
	f "fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"nup/types"
)

type PingService struct {
	config types.Config
	Wg     *sync.WaitGroup
	Exit   chan bool
}

func Init(wg *sync.WaitGroup, config types.Config) *PingService {
	return &PingService{
		config: config,
		Exit:   make(chan bool),
	}
}

func (s *PingService) Ping() {
	client := &http.Client{
		Timeout: time.Duration(s.config.TimeoutSecs) * time.Second,
	}

	for {
		domain := s.config.Domains[rand.Intn(len(s.config.Domains))]
		startTime := time.Now()

		resp, err := client.Get(domain)
		if err != nil {
			if s.config.FlagVerbose {
				f.Printf("%s | Status: Failed\n", domain)
			}
			continue
		}
		defer resp.Body.Close()

		log := f.Sprintf("%s", domain)
		switch {
		case s.config.FlagStatus:
			log = f.Sprintf("%s | Status: %s", log, resp.Status)
		case s.config.FlagLatency:
			elapsed := time.Since(startTime)
			log = f.Sprintf("%s | Time: %ds.%03ds", log, int64(elapsed.Seconds()), elapsed.Milliseconds())
		}

		if s.config.FlagVerbose {
			f.Println(log)
		}

		time.Sleep(time.Duration(s.config.IntervalSecs) * time.Second)

		if <-s.Exit {
			f.Println("Exiting...(ctx: ping.go)")
			break
		}
	}
}
