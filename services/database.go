package services

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseService struct {
	db *sql.DB
}

func (ds *DatabaseService) prepareTables() {
	statement, _ := ds.db.Prepare("CREATE TABLE IF NOT EXISTS log (id INTEGER PRIMARY KEY, log INTEGER)")
	statement2, _ := ds.db.Prepare("CREATE TABLE IF NOT EXISTS config (id INTEGER PRIMARY KEY, key STRING, value STRING)")
	statements := []*sql.Stmt{statement, statement2}
	for _, statement := range statements {
		statement.Exec()
	}
}

func (ds *DatabaseService) Get(query string, args ...interface{}) (*sql.Rows, error) {
	return ds.db.Query(query, args...)
}

func (ds *DatabaseService) Post(query string, args ...interface{}) (sql.Result, error) {
	return ds.db.Exec(query, args...)
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
