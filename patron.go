package main

import "database/sql"

func addPatronToDB(userID string, status uint8) error {
	db, err := initDB()

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO patrons (userID, status) VALUES (?, ?)", userID, status)

	return err
}

func removePatronFromDB(userID string) error {
	db, err := initDB()

	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM patrons WHERE userID = ?", userID)

	return err
}

// from patron.go on the main bot
func getPatronStatus(userID string) (uint8, error) {
	db, err := initDB()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	output, err := db.Prepare("SELECT status FROM patrons WHERE userID = ?")

	if err != nil {
		return 0, err
	}

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

func getAllPatrons() ([]boostedUser, error) {
	db, err := initDB()

	if err != nil {
		return []boostedUser{}, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT userID, status FROM patrons")

	if err != nil {
		return []boostedUser{}, err
	}

	var userID string
	var status uint8
	var dbPatrons []boostedUser
	for rows.Next() {
		err = rows.Scan(&userID, &status)
		if err != nil {
			return []boostedUser{}, err
		}

		patron := boostedUser{
			ID:     userID,
			Status: status,
			Guilds: []string{},
		}
		dbPatrons = append(dbPatrons, patron)
	}

	return dbPatrons, nil
}
