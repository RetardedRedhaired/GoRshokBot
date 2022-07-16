package internal

import "github.com/bwmarrin/discordgo"

func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateGameStatus(0, "Ем мясо, пивом запиваю")
}