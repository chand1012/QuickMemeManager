package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func checkThread(discord *discordgo.Session) {

	fmt.Println("Starting Check Thread.")
	fmt.Println("Generating lock file.")
	testData, err := lockFileCreate()

	if err != nil {
		panic(err)
	}

	fmt.Println("Thread Started.")

	for {

		var wg sync.WaitGroup

		fileEqual, err := lockFileEqu(testData)

		if err != nil {
			fmt.Println(err)
			break
		}

		if !fileEqual {
			fmt.Println("New process thread started, killing old thread.")
			break
		}

		// this will get all the user ids that are patrons, then check the database against each one.
		// There will be one SQL query to get all of the rows into memory, then
		// check via goroutines. If a user's ID is in the database, the bot will do nothing
		// if a user's ID is not in the database, then a function that would handle adding the user would be
		// executed. If a few IDs are found in the database that are not in the patrons role, they would be deleted.

		dbPatrons, err := getAllBoostedUsers()
		if err != nil {
			fmt.Println(err)
			break
		}

		patrons, err := getAllServerBenefactors(discord)
		if err != nil {
			fmt.Println(err)
			break
		}

		// this loop makes sure that all the patrons on the server
		// are in the database. If the are not, add them to it
		for _, patron := range patrons {
			exists := false
			for _, dbPatron := range dbPatrons {
				if patron.ID == dbPatron.ID {
					exists = true
					break
				}
			}
			if exists {
				continue
			} else {
				err = addPatronToDB(patron.ID, patron.Status)
				if err != nil {
					fmt.Println(err)
					break
				}
				err = sendBoostRequest(discord, patron.ID, patron.Status)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}

		// need a loop that deletes people who are no longer patrons from
		// the database

	}
}

// from memeQueue.go of the main bot.
func lockFileEqu(input []byte) (bool, error) {
	data, err := ioutil.ReadFile("./thread.lock")
	if err != nil {
		return false, err
	}
	if bytes.Compare(input, data) == 0 {
		return true, nil
	}
	return false, nil
}

func lockFileExists() bool {
	info, err := os.Stat("./thread.lock")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func lockFileCreate() ([]byte, error) {
	fileData := make([]byte, 8)
	rand.Read(fileData)
	err := ioutil.WriteFile("./thread.lock", fileData, 0644)
	return fileData, err
}
