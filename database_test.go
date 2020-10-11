package main

import (
	"testing"
)

func TestConnection(t *testing.T) {
	db, err := initDB()

	if err != nil {
		t.Errorf("There was an error establishing the database connection: %v", err)
	}

	if db == nil {
		t.Errorf("There was an error connecting to the DB, DB pointer is nil.")
	}
}

func TestPatrons(t *testing.T) {
	var err error
	var status uint8

	testID := "00000000000"

	// These will be changed as deemed necessary.
	// minStatus will probably remain 1
	// maxStatus could change as more roles
	// and tiers are introduced.
	const minStatus = 1
	const maxStatus = 2

	for i := minStatus; i <= maxStatus; i++ {
		err = addPatronToDB(testID, uint8(i))

		if err != nil {
			t.Errorf("There was an error adding Patron to the Database: %v", err)
			break
		}

		status, err = getPatronStatus(testID)

		if err != nil {
			t.Errorf("There was an error getting the Patron from the Database: %v", err)
			break
		}

		if status != uint8(i) {
			t.Errorf("There was an error getting the Patron from the database: expected status %d, got %d.", i, status)
			break
		}

		err = removePatronFromDB(testID)

		if err != nil {
			t.Errorf("There was an error removing the Patron from the database: %v", err)
			break
		}
	}
}

func TestBoosts(t *testing.T) {
	var err error
	const status = 2
	testID := "thisisatest"
	testGuild := "626209936262823937" // This is the ID of the official QuickMeme server

	err = setBoostedUser(testID, status, testGuild)

	if err != nil {
		t.Errorf("There was an error adding Boost to the Database: %v", err)
	}

	users, err := getAllBoostedUsers()

	if err != nil {
		t.Errorf("There was an error getting boosted users: %v", err)
	}

	foundTestUser := false
	for _, user := range users {
		if user.ID == testID {
			foundTestUser = true
			break
		}
	}

	if !foundTestUser {
		t.Errorf("The test user was not found in the Boosted table!")
	}

	err = removeBoostedUser(testID)

	if err != nil {
		t.Errorf("There was an error removing the boosted user: %v", err)
	}
}
