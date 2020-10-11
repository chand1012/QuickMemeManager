package main

import (
	"os"
	"testing"
)

func TestDBEnv(t *testing.T) {
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWD")
	database := os.Getenv("DB")
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")

	endstr := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database
	dbStr := getDBEnv()

	if endstr != dbStr {
		t.Errorf("Got unexpected database string. Expected %s got %s.", endstr, dbStr)
	}
}
