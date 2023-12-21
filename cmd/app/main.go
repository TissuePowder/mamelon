package main

import (
	"encoding/json"
	"os"

	"mamelon/internal/logger"
	"mamelon/internal/models"
	"mamelon/internal/scraper"
)

var LinkSet map[string]struct{}

type application struct {
	config  models.Config
	logger  *logger.Logger
	scraper *scraper.Scraper
}

func main() {

	LinkSet = make(map[string]struct{})

	logger := logger.New()
	scraper := scraper.New()

	data, err := os.ReadFile("config.json")
	if err != nil {
		logger.Error(err.Error())
	}

	var cfg models.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		logger.Error(err.Error())
	}

	app := &application{
		config: cfg,
		logger: logger,
		scraper: scraper,
	}

	err = app.getLinks("tweets.txt")

	if err != nil {
		logger.Error(err.Error())
	}

	err = app.runBot()

	if err != nil {
		logger.Error(err.Error())
	}
}
