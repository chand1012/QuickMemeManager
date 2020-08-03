package main

import "github.com/bwmarrin/discordgo"

// this file will contain the functions that
// get the user data from discord

type benefactor struct {
	ID     string
	Status uint8
}

func getAllServerBenefactors(discord *discordgo.Session) ([]benefactor, error) {
	var benefactors []benefactor

	guildID := "626209936262823937"
	roles := []string{"739935672416206848", "739935562621780028", "739935406082228336"}
	members, err := discord.GuildMembers(guildID, "", 1000)

	if err != nil {
		return nil, err
	}

	for _, member := range members {
		for _, role := range member.Roles {
			if stringInSlice(role, roles) {
				tempUser := benefactor{
					ID: member.User.ID,
				}
				if role == "739935672416206848" || role == "739935562621780028" {
					tempUser.Status = 2
				} else {
					tempUser.Status = 1
				}
				benefactors = append(benefactors, tempUser)
			}
		}
	}

	return benefactors, nil
}
