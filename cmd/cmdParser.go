package main

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

func registerRouter(session *discordgo.Session) {
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Responds with 'pong!'",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		Handler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong!")
		},
	})

	router.Initialize(session)
}
