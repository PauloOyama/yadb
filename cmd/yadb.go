package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/agstrc/yadb/internal/interaction"
	"github.com/agstrc/yadb/internal/util"
	"github.com/bwmarrin/discordgo"
)

func main() {
	var update bool

	flag.BoolVar(&update, "u", false, "If set, updates the bot's registered subcommands and exits")
	flag.Parse()

	if update {
		updateBot()
	} else {
		runBot()
	}
}

// runBot enables the bot's processing of requests
func runBot() {
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

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
	fmt.Println("Closing the websocket...")
}

// upbdateBot deletes all the bot's slash commands and registers the commands set in interactions package
func updateBot() {
	session, err := discordgo.New("Bot " + util.GetEnv("BOT_TOKEN"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	err = session.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	fmt.Println("Deleting current slash commands...")
	deleteSlashCommands(session)
	fmt.Println("Registering slash commands...")
	registerSlashCommands(session)
}
