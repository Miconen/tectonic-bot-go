package commands

import (
	"github.com/bwmarrin/discordgo"
)

var command = &CommandOptions {
    Name: "helloworld",
    Description: "A Simple test command",

    f: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "Hello World!",
            },
        })
    },
}

func init() {
    CommandStorage = append(CommandStorage, command)
}