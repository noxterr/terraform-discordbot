package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/noxterr/terraform-discordbot/commands"
	"github.com/noxterr/terraform-discordbot/src"
)

func main() {
	bot, err := discordgo.New("Bot " + src.Initialize())
	if err != nil {
		panic(err)
	}

	bot.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "ping": // This maps directly to terraform's variable
			commands.Ping(s, i)
		}
	})

	err = bot.Open()
	if err != nil {
		fmt.Println(err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = bot.Close()
	if err != nil {
		fmt.Println(err)
	}
}
