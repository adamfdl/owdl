package main

import (
	"net/http"
	"os"

	"github.com/adamfdl/owly/routers"
	"github.com/adamfdl/owly/task"
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
