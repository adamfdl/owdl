package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/adamfdl/ow_discord_leaderboard/database/redis"
	"github.com/rs/zerolog/log"
)

type Data struct {
	Payload interface{} `json:"data"`
}

type PlayerInfo struct {
	LocalRank       int    `json:"local_rank"`
	Battletag       string `json:"battletag"`
	CompetitiveRank int    `json:"competitive_rank"`
}

func Leaderboard(w http.ResponseWriter, r *http.Request) {

	log.Info().Str("method", "leaderboard").Msg("Endpoint /leaderboard hit")

	result, err := redis.GetOWLeaderboardOperator().RetrieveLeaderboard()
	if err != nil {
		log.Error().
			Str("method", "retrieve_leaderboard_redis").
			Msgf("Failed retrieving leaderboard from redis. Error: %s", err.Error())
		return
	}

	var playerInfos []PlayerInfo
	for i := range result {
		var playerInfo PlayerInfo
		playerInfo.Battletag = strings.Replace(result[i].Member.(string), "-", "#", 1)
		playerInfo.CompetitiveRank = int(result[i].Score)
		playerInfo.LocalRank = i + 1

		log.Debug().
			Int("index", i).
			Float64("rank", result[i].Score).
			Str("battletag", result[i].Member.(string)).
			Msg("Player retrieved from redis")

		playerInfos = append(playerInfos, playerInfo)
	}

	log.Info().Str("method", "leaderboard").Msg("Endpoint /leaderboard finished without errors")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Data{Payload: playerInfos})
}
