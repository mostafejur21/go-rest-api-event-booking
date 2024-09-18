package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mostafejur21/go_rest_api/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/v1/events", getEvents)
	server.GET("/api/v1/events/:id", getEvent)

    // this is how we can use middleware using gin.
    // the Group method is to group which route should be use this middleware
    // and the Use method is to use for using the middleware

    authenticate := server.Group("/")
    authenticate.Use(middleware.Authenticate)
	authenticate.POST("/api/v1/events", createEvents)
	authenticate.PUT("/api/v1/events/:id", updateEvents)
	authenticate.DELETE("/api/v1/events/:id", deleteEvents)

	server.POST("/api/v1/signup", signup)
	server.POST("/api/v1/login", login)
}
