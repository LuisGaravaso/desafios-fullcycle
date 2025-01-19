package usecase

import (
	"fmt"

	"cleanarch/internal/entity"
	"cleanarch/internal/event"
	"cleanarch/pkg/events"
)

type GetOrderUseCase struct {
	OrderUseCase
}

func NewGetOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderUseCase: OrderUseCase{
			OrderRepository: OrderRepository,
			OrderEvent:      event.OrderListed(),
			EventDispatcher: EventDispatcher,
		},
	}
}

func (u *GetOrderUseCase) FindAll() ([]OrderOutputDTO, error) {
	orders, err := u.OrderRepository.GetOrders()
	if err != nil {
		return nil, err
	}

	dto := GetOrderOutputDTO{
		Message:    "All orders fetched",
		OrderCount: len(orders),
	}

	u.OrderEvent.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.OrderEvent)

	//converte orders in OrderOutputDTO
	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			Exists:     true,
		})
	}

	return ordersDTO, nil
}

func (u *GetOrderUseCase) FindById(input GetOrderInputDTO) (OrderOutputDTO, error) {

	//Setting up the DTOs
	var ordersDTO OrderOutputDTO
	messageDTO := GetOrderOutputDTO{
		Message:    fmt.Sprintf("Order %s listed", input.ID),
		OrderCount: 1,
	}

	//Getting the order
	order, err := u.OrderRepository.GetOrderById(input.ID)
	if err != nil {
		messageDTO = GetOrderOutputDTO{
			Message:    fmt.Sprintf("Order %s not found", input.ID),
			OrderCount: 0,
		}
	} else {
		ordersDTO = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			Exists:     true,
		}
	}

	u.OrderEvent.SetPayload(messageDTO)
	u.EventDispatcher.Dispatch(u.OrderEvent)

	return ordersDTO, nil
}
