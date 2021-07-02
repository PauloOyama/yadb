package interaction

import (
	"github.com/agstrc/yadb/internal/dex"
	"github.com/bwmarrin/discordgo"
)

// lofiCommand responds with a Spotify link to a lofi playlist
func lofiCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "Here's a great playlist (totally not biased, btw)\n" +
				"https://open.spotify.com/playlist/4ODfi4RIylpJ4z7qYfVAG5?si=f5f38baa446b4828",
		},
	})
}

// randomManga return a manga from mangaDex API
func randomManga(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embeds, err := dex.GetRandom()
	if err != nil {
		ephemeralReply(s, i.Interaction, internalError)
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Embeds: embeds,
		},
	})
}
