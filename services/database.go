package services

import (
	"nup/types"

	"database/sql"
	f "fmt"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseSvc struct {
	db *sql.DB
}

// Getters and Setters
func (ds *DatabaseSvc) Get(query string, args ...interface{}) (*sql.Rows, error) {
	f.Println(query, "Args -> ", args)
	return ds.db.Query(query, args...)
}

func (ds *DatabaseSvc) Post(query string, args ...interface{}) (sql.Result, error) {
	f.Println(query, "Args -> ", args)
	return ds.db.Exec(query, args...)
}

func (ds *DatabaseSvc) Close() {
	ds.db.Close()
}

func (ds *DatabaseSvc) GetConfig() types.Args {
	rows, err := ds.Get("SELECT key, value FROM config")
	if err != nil {
		panic(err)
	}

	args := types.Args{}
	for rows.Next() {
		var key, value string
		rows.Scan(&key, &value)
		switch key {
		case "Latency":
			b, _ := strconv.ParseBool(value)
			args.Latency = b
		case "Status":
			b, _ := strconv.ParseBool(value)
			args.Status = b
		case "Verbose":
			b, _ := strconv.ParseBool(value)
			args.Verbose = b
		case "LogAll":
			b, _ := strconv.ParseBool(value)
			args.LogAll = b
		case "Domains":
			args.Domains = strings.Split(value, ",")
		}
	}

	return args
}

func (ds *DatabaseSvc) SetConfig(args types.Args) {
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "status", args.Status)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "latency", args.Latency)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "verbose", args.Verbose)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "log_all", args.LogAll)
	for _, domain := range args.Domains {
		ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "domain", domain)
	}
}

/*
 * Setup A New Database File
 */
func (ds *DatabaseSvc) prepareTables() {
	statement, _ := ds.db.Prepare(`CREATE TABLE IF NOT EXISTS config (key STRING PRIMARY KEY, value STRING)`)
	statement2, _ := ds.db.Prepare(`CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, latency INTEGER, status TEXT, domain TEXT, time TIMESTAMP)`)
	statements := []*sql.Stmt{statement, statement2}
	for _, statement := range statements {
		statement.Exec()
	}
}

func DbInit() *DatabaseSvc {
	db, err := sql.Open("sqlite3", "./nup.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)

	ds := &DatabaseSvc{db: db}
	ds.prepareTables()

	return ds
}
