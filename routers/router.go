package routers

import (
	"github.com/adamfdl/ow_discord_leaderboard/controller"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", controller.Ping).Methods("GET")
	router.HandleFunc("/leaderboard", controller.Leaderboard).Methods("GET")
	return router
}
