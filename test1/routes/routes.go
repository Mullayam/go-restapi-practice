package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", GetAllEvents)
	server.POST("/", CreateEvent)
	server.GET("/event/:id", GetEventById)
}
