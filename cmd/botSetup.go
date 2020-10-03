package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func registerHandlers(session *discordgo.Session) {
	// Register your handler here
	// e.g. dg.AddHandler(messageCreate)
}

func setup(token string) {

	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	registerHandlers(session)
	registerRouter(session)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	// TODO: this might not be necessary, not sure what it does
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	// Open the websocket and begin listening.
	err = session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("GoCallBot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = session.Close()
}
