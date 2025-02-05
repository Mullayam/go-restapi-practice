package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCors)
	mux.Get("/", app.Hello)
	mux.Get("/test", app.Home)
	mux.
		Route("/api", func(r chi.Router) {
			mux.Use(app.isAuthenticated)
		})
	return mux
}
