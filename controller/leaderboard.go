package controller

import (
	"fmt"

	"github.com/adamfdl/owdl/database/redis"
	"github.com/bwmarrin/discordgo"
)

func OWDiscordLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!owdl standings" {
		var reply string
		if result, err := redis.GetOWLeaderboardOperator().RetrieveLeaderboard(); err != nil {
			reply = ":shit: Technical difficulties :shit:"
			s.ChannelMessageSend(m.ChannelID, reply)
		} else {
			if len(result) <= 0 {
				reply = ":shit: There's no player to retrieve in the database :shit:"
				s.ChannelMessageSend(m.ChannelID, reply)
				return
			}

			var formattedReply string
			for i := 0; i < len(result); i++ {
				player := fmt.Sprintf("[%d]\t> %s\n\t\t\tSR: %d\n", i+1, result[i].Member.(string), int(result[i].Score))
				formattedReply += player
			}

			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("```%s```", formattedReply))
		}
	}
}
