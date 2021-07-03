package interaction

import (
	"github.com/agstrc/yadb/internal/dex"
	dg "github.com/bwmarrin/discordgo"
)

// lofiCommand responds with a Spotify link to a lofi playlist
func lofiCommand(s *dg.Session, i *dg.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionApplicationCommandResponseData{
			Content: "Here's a great playlist (totally not biased, btw)\n" +
				"https://open.spotify.com/playlist/4ODfi4RIylpJ4z7qYfVAG5?si=f5f38baa446b4828",
		},
	})
}

// randomManga return a manga from mangaDex API
func randomManga(s *dg.Session, i *dg.InteractionCreate) {
	embeds, err := dex.GetRandom()
	if err != nil {
		ephemeralReply(s, i.Interaction, internalError)
		return
	}
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionApplicationCommandResponseData{
			Embeds: embeds,
		},
	})
}

// getManga returns the desired manga
func getManga(s *dg.Session, i *dg.InteractionCreate) {

	embeds, err := dex.GetMangaReader(i.Data.Options[0].StringValue())
	if err != nil {
		ephemeralReply(s, i.Interaction, internalError)
		return
	}
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseChannelMessageWithSource,
		Data: &dg.InteractionApplicationCommandResponseData{
			Embeds: embeds,
		},
	})
}
