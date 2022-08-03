# tectonic-bot-go
Attempt at learning Go by rewriting my tectonic-bot project.

## Config
The only required config values are `BotToken` and `BotPrefix`. The bot does have support for additional values, `GuildID` and `RemoveCommands`. These values can also be passed as command-line flags. Each config value has a flag which will override your specified config values. These flags are, `guild`, `token`, `prefix`, and `rmcmd`.
```json
{
	"BotToken": "DISCORD_API_KEY_HERE",
	"BotPrefix": "!"
}

```
