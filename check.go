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
