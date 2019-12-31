package src

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./public/db/active_catalog.db")
	if err != nil {
		panic(err)
	}
}
