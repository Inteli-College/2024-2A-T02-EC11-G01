package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type LocationCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewLocationCreatedHandler(rabbitMQChannel *amqp.Channel) *LocationCreatedHandler {
	return &LocationCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *LocationCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Location created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"",                 // exchange
		"location.created", // key name
		false,              // mandatory
		false,              // immediate
		msg,                // message to publish
	)
}
