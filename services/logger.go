package services

import (
	"nup/types"
)

func LogSql(db *DatabaseSvc, log types.Log) {
	query := "INSERT INTO logs (time, status, latency, domain) VALUES (?, ?, ?, ?)"
	db.Post(query, log.Time, log.Status, log.Latency, log.Domain)
}
