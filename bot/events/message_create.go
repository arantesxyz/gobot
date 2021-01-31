package events

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
