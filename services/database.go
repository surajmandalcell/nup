package services

import (
	"database/sql"
	"nup/types"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseService struct {
	db *sql.DB
}

// Getters and Setters
func (ds *DatabaseService) Get(query string, args ...interface{}) (*sql.Rows, error) {
	return ds.db.Query(query, args...)
}

func (ds *DatabaseService) Post(query string, args ...interface{}) (sql.Result, error) {
	return ds.db.Exec(query, args...)
}

func (ds *DatabaseService) Close() {
	ds.db.Close()
}

func (ds *DatabaseService) GetConfig() types.Args {
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

func (ds *DatabaseService) SetConfig(args types.Args) {
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "status", args.Status)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "latency", args.Latency)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "verbose", args.Verbose)
	ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "log_all", args.LogAll)
	for _, domain := range args.Domains {
		ds.Post("INSERT INTO config (key, value) VALUES (?, ?)", "domain", domain)
	}
}

/*
 * This does proper due dilligance and gives\
 * a pointer to the DatabaseService
 */
func (ds *DatabaseService) prepareTables() {
	statement, _ := ds.db.Prepare(`CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY,	latency INTEGER, status INTEGER, domain TEXT,	time INTEGER)`)
	statement2, _ := ds.db.Prepare("CREATE TABLE IF NOT EXISTS config (id INTEGER PRIMARY KEY, key STRING, value STRING)")
	statements := []*sql.Stmt{statement, statement2}
	for _, statement := range statements {
		statement.Exec()
	}
}

func DbService() *DatabaseService {
	db, err := sql.Open("sqlite3", "./nup.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)

	ds := &DatabaseService{db: db}
	ds.prepareTables() // Prepare Tables

	return ds
}
