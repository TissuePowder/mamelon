package main

import (
	"encoding/json"
	"os"

	"mamelon/internal/logger"
	"mamelon/internal/models"
	"mamelon/internal/scraper"
	"mamelon/internal/translator"
)

var LinkSet map[string]struct{}

type application struct {
	config     models.Config
	logger     *logger.Logger
	scraper    *scraper.Scraper
	translator translator.Translator
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

	translator, err := translator.New(cfg.Translate)
	if err != nil {
		logger.Error(err.Error())
	}

	app := &application{
		config:     cfg,
		logger:     logger,
		scraper:    scraper,
		translator: translator,
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
