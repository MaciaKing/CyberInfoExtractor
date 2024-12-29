package main

import (
	"CyberInfoExtractor/cmd/workers"
	"CyberInfoExtractor/database"
	"CyberInfoExtractor/models"

	"github.com/gin-gonic/gin"
)

func startBackgroundWorkers() {
	go workers.ReadDataToExtract("/go/src/app/dataToExtract/blacklist_domain.txt")
	go workers.ExtractAllQueue()
}

func main() {
	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()
	rb.CloseRabbitMQ()

	database.Connect()
	database.Migrate()

	startBackgroundWorkers()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run("0.0.0.0:8080")
}
