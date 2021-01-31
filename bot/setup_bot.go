package bot

import (
	"log"

	"github.com/arantesxyz/gobot/bot/events"
	"github.com/bwmarrin/discordgo"
)

func SetupBot(token string) (bot *discordgo.Session) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)

	return
}
