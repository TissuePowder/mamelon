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

	ticker := time.NewTicker(60 * time.Second)

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				tweets := app.getTweets()
				for _, c := range app.config.Bot.Channels {

					for _, tweet := range tweets {
						redirectLink := fmt.Sprintf("https://%s%s", app.config.Scrape.Redirect, tweet.Link)
						discord.ChannelMessageSend(c, redirectLink)
						discord.ChannelMessageSend(c, tweet.Text)
						// err := app.addLink(tweet.Link)
						// if err != nil {
						// 	app.logger.Error(err.Error())
						// }
						time.Sleep(1 * time.Second)
					}

					bMsg, err := os.ReadFile("system.txt")
					if err != nil {
						app.logger.Error(err.Error())
					}
					msg := string(bMsg)
					discord.ChannelMessageSend(c, msg)
				}

				err = os.Truncate("system.txt", 0)
				if err != nil {
					app.logger.Error(err.Error())
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

	trigger1 := fmt.Sprintf("<@%s>", discord.State.User.ID)
	trigger2 := app.config.Bot.Trigger

	if strings.HasPrefix(message.Content, trigger1) || strings.HasPrefix(message.Content, trigger2) {
		discord.ChannelTyping(message.ChannelID)
		prompt := strings.TrimPrefix(message.Content, trigger1)
		prompt = strings.TrimPrefix(prompt, trigger2)
		prompt = strings.TrimSpace(prompt)
		app.logger.Info(fmt.Sprintf("%s|%s: %s", message.Author.Username, message.Author.ID, prompt))
		// reply, err := app.getGptResponse(message.Author.ID, prompt)
		if len(prompt) == 0 {
			prompt = "who is mamelon?"
		}
		reply, err := app.getQAResponse(prompt)
		if err != nil {
			app.logger.Error(err.Error())
			return
		}
		discord.ChannelMessageSend(message.ChannelID, reply)
	}

	tweetTexts := app.getTweetTextsFromMessage(message.Content)
	for _, tt := range tweetTexts {
		// msg := &discordgo.MessageSend{
		// 	Content:         tt,
		// 	AllowedMentions: &discordgo.MessageAllowedMentions{RepliedUser: false},
		// 	Reference:       message.Reference(),
		// }
		// discord.ChannelMessageSendComplex(message.ChannelID, msg)
		discord.ChannelMessageSend(message.ChannelID, tt)
	}

}
