package service

import (
	"context"

	"cleanarch/internal/infra/grpc/pb"
	"cleanarch/internal/usecase"
)

type GetOrderService struct {
	pb.UnimplementedGetOrderServiceServer
	GetOrderUseCase usecase.GetOrderUseCase
}

func NewGetOrderService(GetOrderUseCase usecase.GetOrderUseCase) *GetOrderService {
	return &GetOrderService{
		GetOrderUseCase: GetOrderUseCase,
	}
}

func (s *GetOrderService) GetOrderById(ctx context.Context, req *pb.GetOrderByIdRequest) (*pb.Order, error) {
	input := usecase.GetOrderInputDTO{
		ID: req.Id,
	}

	order, err := s.GetOrderUseCase.FindById(input)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
		Exists:     order.Exists,
	}, nil
}

func (s *GetOrderService) GetAllOrders(ctx context.Context, req *pb.Blank) (*pb.GetOrdersResponse, error) {
	orders, err := s.GetOrderUseCase.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.Order
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &pb.GetOrdersResponse{
		Orders: ordersResponse,
	}, nil
}
