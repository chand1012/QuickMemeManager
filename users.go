package main

import "github.com/bwmarrin/discordgo"

// this file will contain the functions that
// get the user data from discord

func getAllServerBenefactors(discord *discordgo.Session) ([]boostedUser, error) {
	var benefactors []boostedUser

	guildID := "626209936262823937"
	roles := []string{"739935672416206848", "739935562621780028", "739935406082228336"}
	members, err := discord.GuildMembers(guildID, "", 1000)

	if err != nil {
		return nil, err
	}

	for _, member := range members {
		for _, role := range member.Roles {
			if stringInSlice(role, roles) {
				tempUser := boostedUser{
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

func sendBoostRequest(discord *discordgo.Session, userID string, status uint8) error {
	var message string

	message = "Hello, and Thank You for supporting the QuickMeme bot on Patreon! "
	message = message + "I am DiscordQuickMeme's manager, and to get started, I will need to know what server"
	if status == 2 {
		message = message + "s "
	} else {
		message = message + " "
	}
	message = message + "you would like to use your benefits on."
	if status == 2 {
		message = message + " You can benefit up to three servers."
	}
	message = message + " To add a server to your benefits, just type `!benefit` in the server of your choice in any channel QuickMemeBot has access to."
	message = message + " To transfer benefits to another server, you must remove the server from your benefits. You can only change your server every 30 days."
	message = message + " To remove your benefits from the server, just type `!unbenefit`. You can then add another server to your benefits."

	pm, err := discord.UserChannelCreate(userID)

	if err != nil {
		return err
	}

	_, err = discord.ChannelMessageSend(pm.ID, message)

	return err

}
