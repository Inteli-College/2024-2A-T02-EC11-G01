package configs

import (
	"log"
	"os"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func setupRabbitMQChannel() (*amqp.Channel, error) {
	rabbitMQChannel, isSet := os.LookupEnv("RABBITMQ_CHANNEL")
	if !isSet {
		log.Fatalf("RABBITMQ_CHANNEL is not set")
	}

	conn, err := amqp.Dial(rabbitMQChannel)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

var setupRabbitMQChannelOnce = sync.OnceValues(setupRabbitMQChannel)

func SetupRabbitMQChannel() (*amqp.Channel, error) {
	return setupRabbitMQChannelOnce()
}
