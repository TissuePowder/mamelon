package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (app *application) runBot() error {
	BotToken := ""
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(app.messageHandler)

	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				discord.ChannelMessageSend(app.config.ChannelID, "15s interval")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	discord.Open()
	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return nil
}

func (app *application) messageHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "m!help"):
		discord.ChannelMessageSend(message.ChannelID, "TODO help")
	}
}
