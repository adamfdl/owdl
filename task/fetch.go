package task

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/adamfdl/owdl/database/redis"
	"github.com/adamfdl/owdl/provider"
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
}

func asyncOverwatchAPIRequest() {

	log.Info().Str("method", "start_cron").Msg("Overwatch API fetch starts")

	battletags := []string{
		"ADAMS-11710",
		"ADAMMS-1274",
		"Awkstronaut-1974",
		"Yoshinoya-11403",
		"BigTummy13-1379",
		"Alphacuremom-11921",
		"flyingpan-11823",
		"HighCadence-11889",
		"Rinmix-1426",
		"SkepticFrog-1100",
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
	interval, _ := strconv.Atoi(os.Getenv("JOB_INTERVAL"))
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				asyncOverwatchAPIRequest()
			}
		}
	}()
}
