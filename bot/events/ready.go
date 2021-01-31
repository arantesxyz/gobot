package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "with discordgo")

	fmt.Println("Logged in as " + string(s.State.User.ID))
}
