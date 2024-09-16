package configs

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupRabbitMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://consumerUser:consumerPassword@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		defer conn.Close()
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		defer ch.Close()
	}
	return ch
}
