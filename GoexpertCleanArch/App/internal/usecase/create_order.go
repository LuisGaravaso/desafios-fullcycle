package usecase

import (
	"cleanarch/internal/entity"
	"cleanarch/internal/event"
	"cleanarch/pkg/events"
)

type CreateOrderUseCase struct {
	OrderUseCase
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderUseCase: OrderUseCase{
			OrderRepository: OrderRepository,
			OrderEvent:      event.OrderCreated(),
			EventDispatcher: EventDispatcher,
		},
	}
}

func (c *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		Exists:     true,
	}

	c.OrderEvent.SetPayload(dto)
	c.EventDispatcher.Dispatch(c.OrderEvent)

	return dto, nil
}
