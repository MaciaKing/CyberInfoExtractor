package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	rabbitConn    *amqp.Connection
	rabbitChannel *amqp.Channel
}

func (rb *Rabbitmq) PushDataToQueue(qName, dataToSend string) {
	if rb.rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
		return
	}

	q, err := rb.rabbitChannel.QueueDeclare(
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

	err = rb.rabbitChannel.PublishWithContext(ctx,
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
	if rb.rabbitChannel == nil {
		log.Println("RabbitMQ channel is not initialized")
		return nil
	}

	// Declarar la cola
	q, err := rb.rabbitChannel.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Configurar el canal de consumo de mensajes
	messages, err := rb.rabbitChannel.Consume(
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
		defer close(msgChan)
		for msg := range messages {
			// Enviar el mensaje al canal de salida
			msgChan <- string(msg.Body)
		}
	}()

	return nil
}

func (rb *Rabbitmq) InitRabbitMQ() {
	var err error
	connURL := "amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@rabbitmq:5672/"

	// Reconnection attempts
	maxRetries := 20
	retryDelay := 5 * time.Second

	for retries := 0; retries < maxRetries; retries++ {
		// try to connect Rabbit server
		rb.rabbitConn, err = amqp.Dial(connURL)
		if err == nil {
			fmt.Println("Conectado a RabbitMQ con éxito")
			break
		}

		// if connection fails, wait to reconnect
		fmt.Printf("Error al conectar a RabbitMQ (intento %d/%d): %s\n", retries+1, maxRetries, err)
		if retries < maxRetries-1 {
			time.Sleep(retryDelay) // wait before retrying
		} else {
			failOnError(err, "Failed to connect to RabbitMQ after multiple retries")
			return
		}
	}

	// try to open rabbitmq channel
	rb.rabbitChannel, err = rb.rabbitConn.Channel()
	failOnError(err, "Failed to open a channel")
}

// Close connection to Rabbitmq server
func (rb *Rabbitmq) CloseRabbitMQ() {
	if rb.rabbitChannel != nil {
		err := rb.rabbitChannel.Close()
		failOnError(err, "Failed to close RabbitMQ channel")
	}
	if rb.rabbitConn != nil {
		err := rb.rabbitConn.Close()
		failOnError(err, "Failed to close RabbitMQ connection")
	}
	log.Println("RabbitMQ connection closed successfully")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
