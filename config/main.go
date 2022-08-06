package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

var (
    GuildID string
    BotToken string // Stored value of BotToken from config.json
    BotPrefix string // Stored value of BotPrefix from config.json
    RemoveCommands bool

    config *configStruct // Store values from config.json
)

type configStruct struct {
    GuildID string `json:"GuildID"`
    BotToken string `json:"BotToken"`
    BotPrefix string `json:"BotPrefix"`
    RemoveCommands bool `json:"RemoveCommands"`
}

func Init(guild string, token string, prefix string, rmcmd bool) {
    err := read(guild, token, prefix , rmcmd)
    if err != nil {
        log.Fatalf("    Error reading config file: %v", err)
    }
}

func read(guild string, token string, prefix string, rmcmd bool) error {
    log.Println("   Reading config file...")
    file, err := ioutil.ReadFile("./config.json")

    if err != nil {
        log.Fatalf("    Couldn't read config file: %v", err)
        return err
    }

    err = json.Unmarshal(file, &config)
    if err != nil {
        log.Fatalf("    Unmarshaling error: %v", err)
    }

    BotToken, err = strCoalesce(prefix)
    if err != nil {
        BotToken = config.BotToken
    }

    BotPrefix, err = strCoalesce(prefix)
    if err != nil {
        BotPrefix = config.BotPrefix
    }

    GuildID, err = strCoalesce(guild)
    if err != nil {
        GuildID = config.GuildID
    }

    // TODO: Use config.json if flag not specified
    RemoveCommands = rmcmd;

    return nil
}

func strCoalesce(s string) (string, error) {
    empty := ""
    if s == empty {
        return s, errors.New("  Empty string")
    }
    return s, nil

}
