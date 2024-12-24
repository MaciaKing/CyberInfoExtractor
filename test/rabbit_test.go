package test

import (
	"CyberInfoExtractor/cmd/globals"
	"CyberInfoExtractor/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendAndReadDataRabbit(t *testing.T) {
	rb := models.Rabbitmq{}
	rb.InitRabbitMQ()

	var myArray = [5]string{"apple", "banana", "cherry", "date", "fig"}

	for i := 0; i < len(myArray); i++ {
		rb.PushDataToQueue(globals.QueueTest1, myArray[i])
	}
	data_readed := make(chan string)

	err := rb.ReadDataFromQueue(globals.QueueTest1, data_readed)
	assert.NoError(t, err)

	for i := 0; i < len(myArray); i++ {
		read := <-data_readed
		assert.Equal(t, read, myArray[i])
	}

}
