package main

import (
	"database/sql"
	"fmt"
	"time"
)

func getAllBoostedUsers() ([]boostedUser, error) {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return nil, err
	}

	tempMap := make(map[string]boostedUser)

	rows, err := db.Query("SELECT userID, status, guildID FROM boosted")

	var userID string
	var status uint8
	var guildID string

	for rows.Next() {
		err = rows.Scan(&userID, &status, &guildID)
		if err != nil {
			return nil, err
		}

		user, ok := tempMap[userID]

		if ok {
			user.Guilds = append(user.Guilds, guildID)
			tempMap[userID] = user
		} else {
			user = boostedUser{
				ID:     userID,
				Status: status,
				Guilds: []string{guildID},
			}
			tempMap[userID] = user
		}
	}

	var users []boostedUser

	for _, v := range tempMap {
		users = append(users, v)
	}

	return users, nil

}

// Currently only used for testing
func setBoostedUser(userID string, status uint8, guild string) error {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return err
	}

	insert, err := db.Prepare("INSERT INTO boosted (userID, status, guildID, cooldown) VALUES (?, ?, ?, ?)")

	if err != nil {
		fmt.Println(err)
		return err
	}

	cooldown := time.Now().Unix() + 2700000

	_, err = insert.Exec(userID, status, guild, cooldown)
	insert.Close()

	return err
}

func removeBoostedUser(userID string) error {
	db, err := initDB()

	defer db.Close()

	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM boosted WHERE userID = ?", userID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
