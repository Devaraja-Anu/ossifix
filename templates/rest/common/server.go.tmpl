package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *application) server() error {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return  nil
}