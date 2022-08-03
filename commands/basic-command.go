package commands

import (
	"github.com/bwmarrin/discordgo"
)



func basicCommand() CommandOptions {
    c := &CommandOptions {
        Name: "basic-command",
        Description: "Basic command",
        f: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
            s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
                Type: discordgo.InteractionResponseChannelMessageWithSource,
                Data: &discordgo.InteractionResponseData{
                    Content: "Hey there! Congratulations, you just executed your first slash command",
                },
            })
        },
    }

    return *c
}
