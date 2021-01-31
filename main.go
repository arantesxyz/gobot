package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/arantesxyz/gobot/bot"
)

func main() {
	bot := bot.SetupBot(os.Getenv("BOT_TOKEN"))

	if err := bot.Open(); err != nil {
		log.Fatal("Error connecting with discord: ", err)
		return
	}

	defer bot.Close()

	fmt.Printf("Bot started! Type `quit` to stop.\n> ")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		if text := scanner.Text(); text == "quit" {
			return
		}

		fmt.Printf("Invalid command!\n> ")
	}
}
