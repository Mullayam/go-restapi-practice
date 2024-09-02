package routes

import (
	"enjoys.in/first-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllEvents(c *gin.Context) {
	events, err := models.GetALlEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
func GetEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
func CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindBodyWithJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}
	event.ID = 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusOK, gin.H{
		"event":   event,
		"message": "event created successfully",
	})
}
