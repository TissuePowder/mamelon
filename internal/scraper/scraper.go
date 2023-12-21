package scraper

import (
	"mamelon/internal/logger"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	c  *colly.Collector
	mu sync.Mutex
}

func New() *Scraper {

	c := colly.NewCollector()
	logger := logger.New()

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
		logger.Error(e.Error())
	})

	return &Scraper{
		c: c,
	}
}

func (s *Scraper) ScrapePages(urls []string) map[string]string {

	m := make(map[string]string)

	s.c.OnHTML(".timeline-item", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a.tweet-link", "href")
		link = strings.TrimSuffix(link, "#m")
		text := e.ChildText(".tweet-content")
		text = strings.ReplaceAll(text, "nitter.net", "twitter.com")
		text = strings.ReplaceAll(text, "piped.video", "youtu.be")
		text = strings.ReplaceAll(text, "teddit.net", "reddit.com")
		s.mu.Lock()
		m[link] = text
		s.mu.Unlock()
	})

	for _, url := range urls {
		s.c.Visit(url)
	}

	return m

}
