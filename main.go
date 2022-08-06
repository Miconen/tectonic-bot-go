package main

import (
	"flag"
	"log"

	"github.com/Miconen/tectonic-bot-go/bot"
	"github.com/Miconen/tectonic-bot-go/commands"
	"github.com/Miconen/tectonic-bot-go/config"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
    BotPrefix      = flag.String("prefix", "", "Bot command prefix for simple commands")
	RemoveCommands = flag.Bool("rmcmd", false, "Remove all commands after shutdowning or not")
)

func init() {
	log.Println("Initializing...")

	log.Println("Parsing flags...")
    flag.Parse()

	log.Println("Initializing config")
	config.Init(*GuildID, *BotToken, *BotPrefix, *RemoveCommands)

	log.Println("Initializing commands")
    commands.Init(bot.Session)

	log.Println("Initializing bot")
	bot.Init()
}

func main() {
    bot.Start()
}
