package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"

	"mamelon/internal/models"
)

type application struct {
	config models.Config
	logger *slog.Logger
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	data, err := os.ReadFile("config.json")
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}

	var cfg models.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error())
		fmt.Println(trace)
		os.Exit(1)
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	err = app.runBot()

	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
