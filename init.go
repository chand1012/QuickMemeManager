package main

import "os"

// this will contain all functions
// related to init-ing the bot

func getDBEnv() string { // returns the string that the DB can use
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWD")
	database := os.Getenv("DB")
	host := os.Getenv("DBHOST")
	port := "3306"

	endstr := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database

	return endstr
}
