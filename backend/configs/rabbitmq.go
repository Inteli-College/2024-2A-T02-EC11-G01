package configs

import amqp "github.com/rabbitmq/amqp091-go"

func SetupRabbitMQChannel(rabbitMQChannel string) (*amqp.Channel, error) {
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