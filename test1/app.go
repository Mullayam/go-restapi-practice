package main

import (
	// "enjoys.in/first-app/database"
	// "enjoys.in/first-app/routes"
	_ "enjoys.in/first-app/utility"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paresed, _ := template.ParseFiles("./index.html")
		errors := paresed.Execute(w, nil)
		if errors != nil {
			fmt.Fprintf(w, "Error executing")
		}
	})
	http.ListenAndServe(":8080", nil)
	// database.InitConnection()
	// utility.ComputerClient()
	// server := gin.Default()
	// routes.RegisterRoutes(server)
	// server.Run(":8080")
}
