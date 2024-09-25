package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQConsumer struct {
	RabbitMQChannel *amqp.Channel
}

func NewRabbitMQConsumer(channel *amqp.Channel) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		RabbitMQChannel: channel,
	}
}

func (rc *RabbitMQConsumer) Consume(out chan<- amqp.Delivery, queue string) error {
	msgs, err := rc.RabbitMQChannel.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}
