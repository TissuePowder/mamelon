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