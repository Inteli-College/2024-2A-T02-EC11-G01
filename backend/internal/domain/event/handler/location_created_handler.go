package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	"github.com/streadway/amqp"
	"sync"
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

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
