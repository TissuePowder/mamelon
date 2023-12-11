package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type Tweet struct {
	Link string
	Text string
}

func (app *application) scrape() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("accept", "*/*")
		r.Headers.Set("accept-language", "en-US,en;q=0.9")
		r.Headers.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	})

	c.OnError(func(r *colly.Response, e error) {
		app.logger.Error(e.Error())
	})

	tweets := []Tweet{}

	c.OnHTML(".timeline-item", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a.tweet-link", "href")
		text := e.ChildText(".tweet-content")
		tweets = append(tweets, Tweet{
			Link: link,
			Text: text,
		})
		fmt.Println(link)
		
	})

	c.Visit("https://example.com")

}
