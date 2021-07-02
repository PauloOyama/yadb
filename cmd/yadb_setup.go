package main

import (
	"fmt"
	"os"

	"github.com/agstrc/yadb/internal/interaction"
	"github.com/bwmarrin/discordgo"
)

// registerSlashCommands registers all slash commands defined by the interaction package
func registerSlashCommands(session *discordgo.Session) {
	for _, command := range interaction.SlashCommands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, "841085532385968169", &command)

		// if a command fails to register, the bot will not be "complete". Therefore, we panic
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to register a slash command:", err.Error())
			os.Exit(1)
		}
	}
}

// deleteSlashCommands deletes found in the given session
func deleteSlashCommands(session *discordgo.Session) {
	commands, err := session.ApplicationCommands(session.State.User.ID, "")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to acquire commands list during deletion:", err.Error())
		os.Exit(1)
	}

	for _, command := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, "ODU4MTM5NjIzNTA3MDk5NjYw.YNZyzQ.FTK2K3rI_gFxggbtvT9XM09Txdw", command.ID)
		if err != nil {
			fmt.Printf("Failed to delete command %s: %s", command.Name, err.Error())
		}
	}
}

// handleReady prints a message once the bot is ready
func handleReady(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready to go!")
	})
}
