package usecase

import (
	"cleanarch/internal/entity"
	"cleanarch/pkg/events"
)

type OrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderEvent      events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}
