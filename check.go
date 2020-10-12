package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

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
			fmt.Println("There was an error getting Boosted users.")
			fmt.Println(err)
			break
		}

		patrons, err := getAllServerBenefactors(discord)
		if err != nil {
			fmt.Println("There was an error getting server benefactors.")
			fmt.Println(err)
			break
		}

		// this checks if a user is in the patrons and in the database
		// if a user is not in the patrons they get deleted from the database
		wg.Add(1)
		go databaseCheckWorker(dbPatrons, patrons, &wg)

		// this loop makes sure that all the patrons on the server
		// are in the database. If they are not, add them to it
		for _, patron := range patrons {
			exists := false
			for _, dbPatron := range dbPatrons {
				if patron.ID == dbPatron.ID {
					exists = true
					break
				}
			}
			if !exists {
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

		wg.Wait()

		time.Sleep(time.Minute * 5)

	}
}

// this checks if a user is present in the database
// but not present in any of the patron roles
func databaseCheckWorker(dbPatrons []boostedUser, patrons []boostedUser, wg *sync.WaitGroup) {

	defer wg.Done()

	var exists bool
	for _, dbPatron := range dbPatrons {
		exists = false
		for _, patron := range patrons {
			if patron.ID == dbPatron.ID {
				exists = true
				break
			}
		}
		if !exists {
			err := removeBoostedUser(dbPatron.ID)
			if err != nil {
				fmt.Println(err)
				break
			}
			err = removePatronFromDB(dbPatron.ID)
			if err != nil {
				fmt.Println(err)
				break
			}
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
