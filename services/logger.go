package services

import (
	"nup/types"
)

func LogSql(db *DatabaseSvc, log types.Log) {
	query := "INSERT INTO logs (latency, status, domain, time) VALUES (?, ?, ?, ?)"
	db.Post(query, log.Latency, log.Status, log.Domain, log.Time)
}
