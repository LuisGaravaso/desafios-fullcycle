package service

import (
	"context"

	"cleanarch/internal/infra/grpc/pb"
	"cleanarch/internal/usecase"
)

type CreateOrderService struct {
	pb.UnimplementedCreateOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewCreateOrderService(CreateOrderUseCase usecase.CreateOrderUseCase) *CreateOrderService {
	return &CreateOrderService{
		CreateOrderUseCase: CreateOrderUseCase,
	}
}

func (s *CreateOrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	input := usecase.CreateOrderInputDTO{
		ID:    req.Id,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}

	order, err := s.CreateOrderUseCase.Execute(input)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}
