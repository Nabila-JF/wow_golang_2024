package main

import (
	"go-blog/config"
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect database
	config.ConnectDatabase()

	// Register routes
	routes.RegisterRoutes(r)

	// Jalankan server
	r.Run(":8080")
}
