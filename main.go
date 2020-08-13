package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	key := os.Getenv("DISCORD_TOKEN")

	discord, err := discordgo.New("Bot " + key)

	if err != nil {
		panic(err)
	}

	discord.AddHandler(commandHandler)
	discord.AddHandler(readyHandler)
	err = discord.Open()
	if err != nil {
		panic(err)
	}

	defer discord.Close()

	fmt.Println("Bot Manager is starting up!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	fmt.Println("Cleaning up ..")
	if lockFileExists() {
		err = os.Remove("./thread.lock")
		if err != nil {
			panic(err) // its shutting down anyway
		}
	}
	fmt.Println("Cleanly shutdown.")
}

func readyHandler(discord *discordgo.Session, ready *discordgo.Ready) {
	go updateStatus(discord)
	go checkThread(discord)
	fmt.Println("Finished starting.")
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// does nothing at the moment
}
