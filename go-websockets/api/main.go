package main

import (
	"log"
	"net/http"
	"ws/api/internal/handlers"
	"ws/api/internal/routes"
)

// main is the entry point for the command line interface
// to the go-websockets API.
func main() {
	approutes := routes.Web()
	log.Print("Listening  WsChannel")
	go handlers.ListenToWsChannel()
	log.Print("Listening on port 8080")
	_ = http.ListenAndServe(":8080", approutes)
	// cmd := &cobra.Command{
	// if err := cmd.Execute(); err != nil {
	// 	log.Fatal(err)
	// }

}
