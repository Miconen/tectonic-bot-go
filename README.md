# tectonic-bot-go
Attempt at learning Go by rewriting my tectonic-bot project. Currently in a working state but not as fleshed out as this repos tectonic-bot counterpart written in TypeScript.

## Config (required)
**The bot requires you to create a config.json file with the structure specified below.**

The only required config values are `BotToken` and `BotPrefix`. The bot does have support for additional values, `GuildID` and `RemoveCommands`. These values can also be passed as command-line flags as `guild`, `token`, `prefix`, and `rmcmd`.
```json
{
	"BotToken": "DISCORD_API_KEY_HERE",
	"BotPrefix": "!"
}

```

## Command-line flags
Each config value has a flag which **will override your specified config values**. These flags are `guild`, `token`, `prefix`, and `rmcmd`.

Example to set the `rmcmd/RemoveCommands` flag which will remove all commands on shutdown
```
go run main.go -rmcmd
```
