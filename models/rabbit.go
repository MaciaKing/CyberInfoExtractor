package models

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
}

var rabbitConn *amqp.Connection
var rabbitChannel *amqp.Channel

func (rb *Rabbitmq) PushDataToQueue(qName, dataToSend string) {
	if rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
		return
	}

	q, err := rabbitChannel.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = rabbitChannel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(dataToSend),
		})
	failOnError(err, "Failed to publish a message")
}

func (rb *Rabbitmq) ReadDataFromQueue(qName string, msgChan chan<- string) error {
	if rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
		return nil
	}

	// Declarar la cola
	q, err := rabbitChannel.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Configurar el canal de consumo de mensajes
	messages, err := rabbitChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Loop para leer mensajes continuamente
	go func() {
		for msg := range messages {
			// Enviar el mensaje al canal de salida
			msgChan <- string(msg.Body)
		}
	}()

	return nil
}

// Start connection to rabbitmq server
func (rb *Rabbitmq) InitRabbitMQ() {
	var err error
	rabbitConn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	rabbitChannel, err = rabbitConn.Channel()
	failOnError(err, "Failed to open a channel")
}

// Close connection to Rabbitmq server
func (rb *Rabbitmq) CloseRabbitMQ() {
	if rabbitChannel != nil {
		err := rabbitChannel.Close()
		failOnError(err, "Failed to close RabbitMQ channel")
	}
	if rabbitConn != nil {
		err := rabbitConn.Close()
		failOnError(err, "Failed to close RabbitMQ connection")
	}
	log.Println("RabbitMQ connection closed successfully")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
