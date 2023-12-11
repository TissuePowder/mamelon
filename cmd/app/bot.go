package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (app *application) runBot() error {
	discord, err := discordgo.New("Bot " + app.config.Bot.Token)
	if err != nil {
		return err
	}

	discord.AddHandler(app.messageHandler)

	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				tweets := app.getTweets()
				for _, c := range app.config.Bot.Channels {
					for _, tweet := range tweets {
						tweet := strings.TrimSuffix(tweet, "#m")
						redirectLink := fmt.Sprintf("https://%s%s", app.config.Scrape.Redirect, tweet)
						discord.ChannelMessageSend(c, redirectLink)
						time.Sleep(1 * time.Second)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	discord.Open()
	defer discord.Close()
	app.logger.Info("Bot started!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	app.logger.Info(s.String())
	return nil
}

func (app *application) messageHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "m!help"):
		discord.ChannelMessageSend(message.ChannelID, "TODO: Commands will be implemented in future")
	}
}
