package main

import "database/sql"

func addPatronToDB(userID string, status uint8) error {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO patrons (userID, status) VALUES (?, ?)", userID, status)

	return err
}

func removePatronFromDB(userID string) error {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM patrons WHERE userID = ?", userID)

	return err
}

// from patron.go on the main bot
func getPatronStatus(userID string) (uint8, error) {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return 0, err
	}

	output, err := db.Prepare("SELECT (status) FROM patrons WHERE userID = ?")

	defer output.Close()

	var status uint8

	err = output.QueryRow(userID).Scan(&status)

	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return status, nil
}
