package event

import "time"

var EventsToRegister = []string{
	OrderCreated().Name,
	OrderListed().Name,
}

type OrderEvent struct {
	Name    string
	Payload interface{}
}

func OrderCreated() *OrderEvent {
	return &OrderEvent{
		Name: "OrderCreated",
	}
}

func OrderListed() *OrderEvent {
	return &OrderEvent{
		Name: "OrderListed",
	}
}

func (e *OrderEvent) GetName() string {
	return e.Name
}

func (e *OrderEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderEvent) GetDateTime() time.Time {
	return time.Now()
}
