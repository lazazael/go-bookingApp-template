package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lazazael/go-bookingApp-template/pkg/config"
	"github.com/lazazael/go-bookingApp-template/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.CleanPath)

	mux.Use(NoSurf)
	mux.Use(SessionStateLoad)

	//mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
