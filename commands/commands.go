package commands

import (
	dc "github.com/g33kidd/n00b/discord"
)

// RegisterTwitchCommands registers the twitch command group
func RegisterTwitchCommands(bot *dc.Bot) {
	twitchEdit := dc.NewCommand("twitchedit", "Edits twitch channel info!", TwitchChannelEditCommand)
	twitchEdit.AddParameter("game", "The new game for twitch channel.", true)
	twitchEdit.AddParameter("title", "The new title for twitch channel.", true)

	channel := dc.NewCommand("twitch", "Gets information for a twitch channel.", TwitchChannelInfoCommand)
	channel.AddParameter("channel", "The channel to get info for.", true)

	// twitchGroup := dc.NewCommandGroup("twitch", "Does stuff")
	// twitchGroup.AddCommand(twitchEdit)

	bot.CmdHandler.AddCommand(twitchEdit)
	bot.CmdHandler.AddCommand(channel)
}

// RegisterRandomCommands registers the random command group
func RegisterRandomCommands(bot *dc.Bot) {
	ping := dc.NewCommand("ping", "Ping!", PingPongCommand)
	pong := dc.NewCommand("pong", "Pong!", PingPongCommand)

	randomCat := dc.NewCommand("cat", "Random cat anyone?", RandomCatCommand)

	api := dc.NewCommand("api", "Makes a GET request to a JSON API and shows the content.", APICommand)
	api.AddParameter("url", "The URL to make a request to.", true)

	bot.CmdHandler.AddCommand(api)
	bot.CmdHandler.AddCommand(randomCat)
	bot.CmdHandler.AddCommand(ping)
	bot.CmdHandler.AddCommand(pong)
}

// RegisterMusicCommands register the music command group
func RegisterMusicCommands(bot *dc.Bot) {

