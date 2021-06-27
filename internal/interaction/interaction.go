// Package interaction defines functions and structures needed for processing an "Interaction Create" event
package interaction

import "github.com/bwmarrin/discordgo"

var SlashCommands = []discordgo.ApplicationCommand{
	{
		Name:        "lofi",
		Description: "A good lofi playlist (according to me)",
	},
}

// commandHandlers maps a command name to an appropriate handler
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"lofi": lofiCommand,
}

// InteractionHandler handles "Interaction Create" events. It only recognizes events listed on the SlashCommands
// variable
func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, ok := commandHandlers[i.Data.Name]; ok {
		handler(s, i)
	}
}
