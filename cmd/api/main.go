package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/Nico2220/billease/internal/aws"
	"github.com/Nico2220/billease/internal/data"
)

const (
	version = "1.0.0"
)

type responseFormat map[string]any

type config struct {
	port int
	env  string
	aws  struct {
		bucketName string
	}
}

type application struct {
	config     config
	logger     *slog.Logger
	models     data.Models
	awsService *aws.AWSService
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 3000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "API environment")
	flag.StringVar(&cfg.aws.bucketName, "bucketName", os.Getenv("BILLEASE_BUCKET_NAME"), "AWS bucket name")

	flag.Parse()

	app := &application{
		config:     cfg,
		logger:     slog.New(slog.NewTextHandler(os.Stdout, nil)),
		models:     data.NewModels(),
		awsService: aws.New(cfg.aws.bucketName),
	}

	if err := app.startServer(); err != nil {
		os.Exit(1)
	}

	fmt.Println(app)

}
