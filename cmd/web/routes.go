package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/seemsod1/api-project/internal/config"
	"github.com/seemsod1/api-project/internal/handlers"
	"net/http"
)

// routes sets up the routes for the application
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(EnableCORS)

	mux.Get("/rate", handlers.Repo.Rate)
	mux.Post("/subscribe", handlers.Repo.Subscribe)

	return mux
}
