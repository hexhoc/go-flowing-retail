package service

import (
	"context"

	"github.com/hexhoc/order-service/internal/pb"
	"github.com/hexhoc/order-service/internal/repository"
)

type OrderInterface interface {
	FindAll(ctx context.Context, productRequest *pb.FindAllRequest) (*pb.FindAllResponse, error)
	FindById(ctx context.Context, productRequest *pb.FindByIdRequest) (*pb.FindByIdResponse, error)
	Save(ctx context.Context, productRequest *pb.SaveRequest) (*pb.StatusResponse, error)
	Update(ctx context.Context, productRequest *pb.UpdateRequest) (*pb.StatusResponse, error)
	Delete(ctx context.Context, productRequest *pb.DeleteRequest) (*pb.StatusResponse, error)
}

type OrderService struct {
	orderRepository repository.OrderInterface
}

func NewOrderService(orderRepository repository.OrderInterface) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}
