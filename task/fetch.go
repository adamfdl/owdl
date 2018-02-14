package task

import (
	"fmt"
	"sync"
	"time"

	"github.com/adamfdl/owly/database/redis"
	"github.com/adamfdl/owly/provider"
	"github.com/rs/zerolog/log"
)

var owAPI = &provider.OverwatchAPI{}

func makeRequest(battletag string, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := owAPI.GetProfile(battletag)
	if err != nil {
		log.Error().
			Str("method", "get_profile_ow_api").
			Msgf("Failed to retrieve OW API. Error: %s", err.Error())
		return
	}

	redis.GetOWLeaderboardOperator().ZAdd(battletag, result.Competitive.Rank)

	log.Debug().
		Str("method", "get_profile_ow_api").
		Int("comp_rank", result.Competitive.Rank).
		Int("comp_win", result.Games.Competitive.Won).
		Int("comp_lost", result.Games.Competitive.Lost).
		Str("username", result.Username).
		Msgf("Fetched username: %s", result.Username)
}

func asyncOverwatchAPIRequest() {

	log.Info().Str("method", "start_cron").Msg("Overwatch API fetch starts")

	battletags := []string{
		"ADAMS-11710",
		"ADAMMS-1274",
		"Awkstronaut-1974",
		"Yoshinoya-11403",
	}

	var wg sync.WaitGroup
	wg.Add(len(battletags))

	start := time.Now()
	for _, battletag := range battletags {
		go makeRequest(battletag, &wg)
	}

	wg.Wait()

	log.Info().
		Str("method", "finish_cron").
		Bool("success", true).
		Str("elapsed", fmt.Sprintf("%.2fs elapsed time", time.Since(start).Seconds())).
		Msg("Overwatch API fetch logs success")
}

func FetchOverwatchAPIJob() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				asyncOverwatchAPIRequest()
			}
		}
	}()
}
