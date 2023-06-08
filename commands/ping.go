package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {

	fmt.Println(i.Data.Type().String())
	err := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello world!",
			},
		},
	)

	if err != nil {
		fmt.Println("Error sending message")
		fmt.Println(err)
	}
}
