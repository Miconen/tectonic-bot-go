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

    // Temporary array for commands to add themselves in and where we will fetch them from
    CommandStorage []*CommandOptions
)

type CommandOptions struct {
    Name string
    Description string
    f func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func AddCommand(command *CommandOptions) {
    ListEntry := &discordgo.ApplicationCommand{
        Name: command.Name,
        Description: command.Description,
    }
    List = append(List, ListEntry)

    Handlers[command.Name] = command.f

    log.Println("   Added command: " + command.Name)
}

func AddCommands() {
    for i := range CommandStorage {
        AddCommand(CommandStorage[i])
    } 
}

func Init(s *discordgo.Session) {
    // Handlers for registered commands
    Handlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
    AddCommands()
}
