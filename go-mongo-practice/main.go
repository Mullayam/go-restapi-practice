package main

import (
	"net/http"

	"enjoys.in/mongo-go-test/config"
	"enjoys.in/mongo-go-test/routes"
)

func main() {
	var s = config.GetSession()
	defer s.Close()

	r := routes.InitRoutes()
	routes.Handler(s)
	http.ListenAndServe(":8080", r)
}
