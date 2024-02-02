package main

import (
	"fmt"
	"slices"
	"strings"
)

type Tweet struct {
	Link string
	Text string
}

func (app *application) getTweets() []Tweet {

	urls := []string{}

	for _, user := range app.config.Scrape.Usernames {
		url := fmt.Sprintf("https://%s/%s", app.config.Scrape.Instances[0], user)
		urls = append(urls, url)
	}

	app.mu.Lock()
	m := app.scraper.ScrapePages(urls)
	app.mu.Unlock()

	tweets := []Tweet{}

	for link, text := range m {
		if !app.existsLink(link) && !strings.HasPrefix(text, "RT @") {
			err := app.addLink(link)
			if err != nil {
				app.logger.Error(err.Error())
			}
			translated := app.getTranslation(text, "JA")
			tweets = append(tweets, Tweet{
				Link: link,
				Text: translated,
			})

		}
	}

	slices.Reverse(tweets)
	return tweets

}
