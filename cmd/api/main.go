package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

const (
	version = "1.0.0"
)

type responseFormat map[string]any

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 3000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "API environment")

	flag.Parse()

	app := &application{
		config: cfg,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	if err := app.startServer(); err != nil {
		os.Exit(1)
	}

	fmt.Println(app)

}
