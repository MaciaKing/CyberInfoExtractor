package test

import (
	"CyberInfoExtractor/cmd/globals"
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendAndReadDataRabbit(t *testing.T) {
	models.InitRabbitMQ()
	var myArray = [5]string{"apple", "banana", "cherry", "date", "fig"}

	for i := 0; i < len(myArray); i++ {
		models.PushDataToQueue(globals.DataExtractedQueue, myArray[i])
	}
	data_readed := make(chan string)

	err := models.ReadDataFromQueue(globals.DataExtractedQueue, data_readed)
	assert.NoError(t, err)

	for i := 0; i < len(myArray); i++ {
		read := <-data_readed
		assert.Equal(t, read, myArray[i])
	}

}
