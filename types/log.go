package types

import "time"

type Log struct {
	Latency int64
	Status  string
	Domain  string
	Time    time.Time
}
