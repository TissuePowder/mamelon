package main

import (
	"fmt"
	"slices"
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

	m := make(map[string]string)
	m = app.scraper.ScrapePages(urls)

	tweets := []Tweet{}

	for link, text := range m {
		if !app.existsLink(link) {
			err := app.addLink(link)
			if err != nil {
				app.logger.Error(err.Error())
			}
			translated, err := app.translate(text)
			if err != nil {
				app.logger.Error(err.Error())
				translated = "`Some error occured during DeepL translation`"
			} else {
				translated = fmt.Sprintf("```\nTranslation (DeepL):\n------------------------\n%s\n```", translated)
			}
			tweets = append(tweets, Tweet{
				Link: link,
				Text: translated,
			})

		}
	}

	slices.Reverse(tweets)
	return tweets

}
