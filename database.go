package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// this is the file that contains all database
// related functions

type boostedUser struct {
	ID     string
	Status uint8 // if 1, then user can have one server. If 2, user can have three servers
	Guilds []string
}

func initDB() (*sql.DB, error) {
	connectionStr := getDBEnv()
	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	return db, err
}
