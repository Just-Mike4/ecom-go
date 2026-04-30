package main

import (
	"os"
	"log/slog"
)
func main() {
	cfg := config{
		addr: ":8081",
		db: dbConfig{},
	}

	api:= application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil{
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}

}