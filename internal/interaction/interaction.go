// Package interaction defines functions and structures needed for processing an "Interaction Create" event
package interaction

import "github.com/bwmarrin/discordgo"

const internalError = "An error has occurred."

var SlashCommands = []discordgo.ApplicationCommand{
	{
		Name:        "lofi",
		Description: "A good lofi playlist (according to me)",
	},
	{
		Name:        "random-manga",
		Description: "Returns a random manga",
	},
}

// commandHandlers maps a command name to an appropriate handler
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"lofi":         lofiCommand,
	"random-manga": randomManga,
}

// InteractionHandler handles "Interaction Create" events. It only recognizes events listed on the SlashCommands
// variable
func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, ok := commandHandlers[i.Data.Name]; ok {
		handler(s, i)
	}
}

// ephemeralReply sends a reply to the given interaction which may only be seen by the user who invoked the interaction
func ephemeralReply(s *discordgo.Session, i *discordgo.Interaction, msg string) error {
	return s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: msg,
			Flags:   64,
		},
	})
}
