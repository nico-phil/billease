package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	port int
	env  string
}

type application struct {
	cfg config
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 3000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "API environment")

	app := &application{
		cfg: cfg,
	}

	if err := app.startServer(); err != nil {
		os.Exit(1)
	}

	fmt.Println(app)

}
