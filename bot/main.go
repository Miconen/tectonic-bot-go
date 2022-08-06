package bot

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Miconen/tectonic-bot-go/commands"
	"github.com/Miconen/tectonic-bot-go/config"

	"github.com/bwmarrin/discordgo"
)

// Bot session data
var (
    // Discordgo Session
    Session *discordgo.Session

    // Bot id, name and discriminator
    BotID string
    BotUsername string
    BotDiscriminator string

    //Contains guild-specific data in a string map, where key = guild ID
	// guildData = make(map[string]*GuildData)

	//Contains guild-specific settings in a string map, where key = guild ID
	// guildSettings = make(map[string]*GuildSettings)

	//Contains user-specific settings in a string map, where key = user ID
	// userSettings = make(map[string]*UserSettings)

    // Contains a pointer to the current log file
	logFile *os.File

	//Contains the current uptime
	uptime time.Time

	//Whether or not discordReady() has been called
	isReady bool
)

func Init() {
	var err error
	Session, err = discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalf("	Invalid bot parameters: %v", err)
	}
    
    u, err := Session.User("@me")
    if err != nil {
        log.Fatalf("	Failed to make bot a user: %v", err)
        return
    }

    BotID = u.ID
    BotUsername = u.Username
    BotDiscriminator = u.Discriminator

	Session.AddHandler(func(session *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
			h(session, i)
		}
		log.Println("Interaction created")
	})

	Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		
		log.Printf("	Logged in as: %v#%v", BotUsername, BotDiscriminator)
	})
}

func Start() {
	err := Session.Open()
	if err != nil {
		log.Fatalf("	Cannot open the session: %v", err)
	}

	log.Println("	Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.List))
	for i, v := range commands.List {
		cmd, err := Session.ApplicationCommandCreate(BotID, config.GuildID, v)
		if err != nil {
			log.Panicf("		Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer Session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	if config.RemoveCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := Session.ApplicationCommandDelete(BotID, config.GuildID, v.ID)
			if err != nil {
				log.Panicf("	Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
