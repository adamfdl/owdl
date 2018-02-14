package main

import (
	"net/http"
	"os"

	"github.com/adamfdl/ow_discord_leaderboard/routers"
	"github.com/adamfdl/ow_discord_leaderboard/task"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {

	task.FetchOverwatchAPIJob()

	http.ListenAndServe(":3009", routers.Init())
}
