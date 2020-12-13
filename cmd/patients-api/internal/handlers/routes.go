package handlers

import (
	"log"
	"net/http"

	"github.com/brudnak/myndshft/internal/platform/web"
	"github.com/jmoiron/sqlx"
)

// Routes creates a handler that knows about all of the API routes.
func Routes(logger *log.Logger, db *sqlx.DB) http.Handler {

	app := web.NewApp(logger)

	patients := Patient{
		DB:  db,
		Log: logger,
	}

	app.Handle(http.MethodGet, "/v1/patients", patients.List)
	app.Handle(http.MethodPost, "/v1/patients", patients.Create)
	app.Handle(http.MethodGet, "/v1/patients/{id}", patients.Retrieve)

	return app
}
