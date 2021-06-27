package main

import (
	"log"

	"github.com/agstrc/yadb/internal/interaction"
	"github.com/bwmarrin/discordgo"
)

// registerSlashCommands registers all slash commands defined by the interaction package
func registerSlashCommands(session *discordgo.Session) {
	for _, command := range interaction.SlashCommands {
		_, err := session.ApplicationCommandCreate(session.State.User.ID, "", &command)

		// if a command fails to register, the bot will not be "complete". Therefore, we panic
		if err != nil {
			log.Panicln("Failed to register a slash command:", err.Error())
		}
	}
}

// deleteSlashCommands deletes found in the given session
func deleteSlashCommands(session *discordgo.Session) {
	commands, err := session.ApplicationCommands(session.State.User.ID, "")
	if err != nil {
		log.Panicln("Failed to delete slash commands:", err.Error())
	}

	for _, command := range commands {
		err := session.ApplicationCommandDelete(session.State.User.ID, "", command.ID)
		if err != nil {
			log.Printf("Failed to delete command %s: %s", command.Name, err.Error())
		}
	}
}

// handleReady prints a message once the bot is ready
func handleReady(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is ready!")
	})
}
