package main

import (
	"CyberInfoExtractor/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("0.0.0.0:8080")
}
