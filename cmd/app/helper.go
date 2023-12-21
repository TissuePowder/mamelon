package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func (app *application) existsLink(link string) bool {
	_, ok := LinkSet[link]
	return ok
}

func (app *application) addLink(link string) error {

	LinkSet[link] = struct{}{}

	file, err := os.OpenFile("tweets.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(link + "\n")
	if err != nil {
		return err
	}

	err = writer.Flush()
	return err
}

func (app *application) getLinks(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		LinkSet[line] = struct{}{}
	}

	err = scanner.Err()
	return err
}

func (app *application) getTweetLinksFromMessage(text string) []string {

	pattern := `https://(?:\S*twitter\.com|\S*x\.com)/\S+/status/\S+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)

	for i, match := range matches {
		parts := strings.Split(match, "/")
		parts[2] = app.config.Scrape.Instances[0]
		match = strings.Join(parts, "/")
		matches[i] = match
	}
	// fmt.Println(matches)
	return matches

}

func (app *application) getTweetTextsFromMessage(msg string) []string {
	tweetLinks := app.getTweetLinksFromMessage(msg)

	texts := app.scraper.ScrapeTweets(tweetLinks)
	return texts

}
