package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (

    // List of commands to register
	List []*discordgo.ApplicationCommand
    // Handlers for registered commands
    Handlers map[string]func(session *discordgo.Session, i *discordgo.InteractionCreate)

)

type CommandOptions struct {
    Name string
    Description string
    f func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func RegisterCommand(co CommandOptions) {
    c := &discordgo.ApplicationCommand{
        Name: co.Name,
        Description: co.Description,
    }
    List = append(List, c)

    Handlers[co.Name] = co.f

    log.Println("Added command: " + co.Name)
}

func Init(s *discordgo.Session) {
    // Handlers for registered commands
    Handlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
    RegisterCommand(basicCommand())
}