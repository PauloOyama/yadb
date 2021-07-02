// Package interaction defines functions and structures needed for processing an "Interaction Create" event
package interaction

import dg "github.com/bwmarrin/discordgo"

const internalError = "An error has occurred."

var SlashCommands = []dg.ApplicationCommand{
	{
		Name:        "lofi",
		Description: "A good lofi playlist (according to me)",
	},
	{
		Name:        "random-manga",
		Description: "Returns a random manga",
	},
	{
		Name:        "get-manga",
		Description: "Return the desired manga",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "string-option",
				Description: "String option",
				Required:    true,
			},
		},
	},
}

// commandHandlers maps a command name to an appropriate handler
var commandHandlers = map[string]func(s *dg.Session, i *dg.InteractionCreate){
	"lofi":         lofiCommand,
	"random-manga": randomManga,
	"get-manga":    getManga,
}

// InteractionHandler handles "Interaction Create" events. It only recognizes events listed on the SlashCommands
// variable
func InteractionHandler(s *dg.Session, i *dg.InteractionCreate) {
	if handler, ok := commandHandlers[i.Data.Name]; ok {
		handler(s, i)
	}
}

// ephemeralReply sends a reply to the given interaction which may only be seen by the user who invoked the interaction
func ephemeralReply(s *dg.Session, i *dg.Interaction, msg string) error {
	return s.InteractionRespond(i, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionApplicationCommandResponseData{
			Content: msg,
			Flags:   64,
		},
	})
}
