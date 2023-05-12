package services

import (
	f "fmt"
	"nup/types"
)

type Logger struct {
}

func (l *Logger) sqlLog(db *DatabaseService, log types.Log) {
	_log := f.Sprintf("INSERT INTO logs (latency, status, domain, time) VALUES (%d, %d, '%s', %d)", log.Latency, log.Status, log.Domain, log.Time)
	db.Post(_log)
}
