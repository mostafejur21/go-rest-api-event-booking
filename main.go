package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mostafejur21/go_rest_api/db"
	"github.com/mostafejur21/go_rest_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
