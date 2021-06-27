package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/agstrc/yadb/internal/interaction"
	"github.com/agstrc/yadb/internal/util"
	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := discordgo.New("Bot " + util.GetEnv("BOT_TOKEN"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	handleReady(session)
	session.AddHandler(interaction.InteractionHandler)

	err = session.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	defer session.Close()

	registerSlashCommands(session)
	defer deleteSlashCommands(session)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
	log.Println("Closing websocket session and removing all slash commands...")
}
