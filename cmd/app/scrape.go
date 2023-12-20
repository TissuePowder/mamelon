package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Tweet struct {
	Link string
	Text string
}

func (app *application) getTweets() []Tweet {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	headers := map[string]string{
		"accept":          "*/*",
		"accept-language": "en-US,en;q=0.9",
		"user-agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	c.OnRequest(func(r *colly.Request) {
		for key, value := range headers {
			r.Headers.Set(key, value)
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		app.logger.Error(e.Error())
	})

	tweets := []Tweet{}

	c.OnHTML(".timeline-item", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a.tweet-link", "href")
		link = strings.TrimSuffix(link, "#m")
		if !app.existsLink(link) {
			// LinkSet[link] = struct{}{}
			err := app.addLink(link)
			if err != nil {
				app.logger.Error(err.Error())
			}
			text := e.ChildText(".tweet-content")
			text = strings.ReplaceAll(text, "nitter.net", "twitter.com")
			text = strings.ReplaceAll(text, "piped.video", "youtu.be")
			text = strings.ReplaceAll(text, "teddit.net", "reddit.com")
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
	})

	for _, user := range app.config.Scrape.Usernames {
		url := fmt.Sprintf("https://%s/%s", app.config.Scrape.Instances[0], user)
		c.Visit(url)
	}

	slices.Reverse(tweets)
	return tweets
}
