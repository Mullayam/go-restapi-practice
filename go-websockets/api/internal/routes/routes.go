package routes

import (
	"net/http"
	"ws/api/internal/handlers"

	"github.com/bmizerany/pat"
)

func Web() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Hello))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))
	mux.Get("/home", http.HandlerFunc(handlers.Home))

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Get("/static", http.StripPrefix("/static", fileServer))

	return mux

}
