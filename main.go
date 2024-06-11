package main

import (
	"github.com/gin-gonic/gin"
	"interview-api/config"
	"interview-api/routes"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
