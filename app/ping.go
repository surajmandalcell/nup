package app

import (
	f "fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"nup/services"
	"nup/types"
)

type PingSvc struct {
	config types.Config
	db     *services.DatabaseSvc
	Wg     *sync.WaitGroup
	Exit   chan bool
}

func Init(wg *sync.WaitGroup, db *services.DatabaseSvc, config types.Config) *PingSvc {
	return &PingSvc{
		config: config,
		db:     db,
		Exit:   make(chan bool),
	}
}

func (s *PingSvc) Ping() {
	domain := s.config.Domains[rand.Intn(len(s.config.Domains))]
	startTime := time.Now()
	elapsed := time.Since(startTime)

	client := &http.Client{
		Timeout: time.Duration(s.config.TimeoutSecs) * time.Second,
	}

	for {
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
			log = f.Sprintf("%s | Time: %ds.%03ds", log, int64(elapsed.Seconds()), elapsed.Milliseconds())
		}

		if s.config.FlagVerbose {
			f.Println(log)
		}

		services.LogSql(s.db, types.Log{
			Latency: int64(elapsed.Milliseconds()),
			Status:  resp.Status,
			Domain:  domain,
			Time:    time.Now(),
		})

		select {
		case <-s.Exit:
			f.Println("Exiting...(ctx: ping.go)")
			s.Wg.Done()
			break
		case <-time.After(0):
			time.Sleep(time.Duration(s.config.IntervalSecs) * time.Second)
		}
	}
}
