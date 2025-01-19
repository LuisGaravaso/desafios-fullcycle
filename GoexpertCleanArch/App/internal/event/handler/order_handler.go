package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"cleanarch/pkg/events"

	"github.com/streadway/amqp"
)

type OrderHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderHandler(rabbitMQChannel *amqp.Channel) *OrderHandler {
	return &OrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: %v \n", event.GetName(), event.GetPayload())
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
