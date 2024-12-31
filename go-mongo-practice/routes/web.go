package routes

import (
	"enjoys.in/mongo-go-test/controllers"
	middleware "enjoys.in/mongo-go-test/middleware"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type Controller struct {
	Session *mgo.Session
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.EnableCors) //attach JWT auth middleware
	// r.Use(middleware.JwtAuthMiddleware) //attach JWT auth middleware
	r.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user", controllers.AddUsers).Methods("POST")
	r.HandleFunc("/user", controllers.UpdateUsers).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/*", controllers.GetAllUsers).Methods("ALL")

	return r
}
func Handler(s *mgo.Session) *Controller {
	return &Controller{s}
}
