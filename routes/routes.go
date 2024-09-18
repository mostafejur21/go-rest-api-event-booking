package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/v1/events", getEvents)
	server.GET("/api/v1/events/:id", getEvent)
	server.POST("/api/v1/events", createEvents)
	server.PUT("/api/v1/events/:id", updateEvents)
	server.DELETE("/api/v1/events/:id", deleteEvents)
	server.POST("/api/v1/signup", signup)
	server.POST("/api/v1/login", login)
}
