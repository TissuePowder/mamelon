package main

import (
	"bufio"
	"os"
)

func (app *application) existsLink(link string) bool {
	_, ok := LinkSet[link]
	return ok
}

func (app *application) addLink(link string) error {

	LinkSet[link] = struct{}{}

	file, err := os.OpenFile("tweets.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
