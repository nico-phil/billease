package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) startServer() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Print("server string at:")

	err := srv.ListenAndServe()
	return err
}
