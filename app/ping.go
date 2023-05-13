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
	exit   chan bool
}

func Init(wg *sync.WaitGroup, db *services.DatabaseSvc, config types.Config) *PingSvc {
	return &PingSvc{
		config: config,
		db:     db,
		Wg:     wg,
		exit:   make(chan bool),
	}
}

func (s *PingSvc) Stop() {
	s.exit <- true
	s.Wg.Done()
}

func (s *PingSvc) Ping() bool {
	domain := s.config.Domains[rand.Intn(len(s.config.Domains))]

	client := &http.Client{
		Timeout: time.Duration(s.config.TimeoutSecs) * time.Second,
	}

	for {
		startTime := time.Now()
		resp, err := client.Get(domain)
		if err != nil {
			if s.config.FlagVerbose {
				f.Printf("%s | Status: Failed\n", domain)
			}
			continue
		}
		resp.Body.Close()
		elapsed := time.Since(startTime)

		log := domain
		switch {
		case s.config.FlagStatus:
			log = f.Sprintf("%s | Status: %s", log, resp.Status)
		case s.config.FlagLatency:
			log = f.Sprintf("%s | Time: %v", log, elapsed.Round(time.Millisecond).String())
		}

		if s.config.FlagVerbose {
			f.Println(log)
		}

		services.LogSql(s.db, types.Log{
			Latency: elapsed.Milliseconds(),
			Status:  resp.Status,
			Domain:  domain,
			Time:    time.Now(),
		})

		select {
		case <-s.exit:
			goto END
		case <-time.After(0):
			time.Sleep(time.Duration(s.config.IntervalSecs) * time.Second)
		}
	}
END:
	return true
}
