package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PredictionCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewPredictionCreatedHandler(rabbitMQChannel *amqp.Channel) *PredictionCreatedHandler {
	return &PredictionCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *PredictionCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Prediction created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",         // exchange
		"prediciton.created", // key name
		false,                // mandatory
		false,                // immediate
		msg,                  // message to publish
	)
}
