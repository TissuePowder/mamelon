package main

import (
	"log/slog"
	"os"
	"runtime/debug"
)

type config struct {
	ChannelID string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	cfg := &config{
		ChannelID: "746435944147714080",
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	app := &application{
		config: *cfg,
		logger: logger,
	}

	err := app.runBot()

	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}
