package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/adamfdl/owdl/controller"
	"github.com/adamfdl/owdl/task"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Error().
			Msgf("Failed creating discord session. Error: %s", err.Error())
		return
	}
	startDiscord(dg)

	task.FetchOverwatchAPIJob()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func startDiscord(dg *discordgo.Session) {
	err := dg.Open()
	if err != nil {
		log.Error().
			Msgf("Failed to opn connection. Error: %s ", err.Error())
		return
	}

	log.Info().
		Msg("Success. OWDL bot is now running. Ctrl-C to exit.")

	dg.AddHandler(controller.OWDiscordLeaderboard)
}
