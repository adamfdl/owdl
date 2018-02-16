package controller

import (
	"fmt"
	"os"
	"time"
)

var separator = "--------------------------------\n"

func buildLeaderboardMessage(leaderboard string) string {
	header := "📈 Rank | Name\n\n"
	return fmt.Sprintf(":city_dusk:  **Overwatch Discord League for Gibraltar Fuccbois**\n\n```xl\n%s%s%s%s```\n", header, leaderboard, separator, seasonEnds())
}

func seasonEnds() string {
	seasonEnds, err := time.Parse("1/2/06", os.Getenv("SEASON_ENDS"))
	if err != nil {
		fmt.Println(err.Error())
	}
	remainingHours := int(time.Until(seasonEnds).Hours())

	if remainingHours <= 24 {
		return fmt.Sprintf("Season ends in %d hours!", remainingHours)
	}

	return fmt.Sprintf("Season ends in %d days!", remainingHours/24)
}
