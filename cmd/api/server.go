package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) startServer() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	app.logger.Info("starting server", "addr", app.config.port, "env", app.config.env)

	err := srv.ListenAndServe()
	return err
}
