package service

import (
	"context"
	"math/big"
	"time"

	"github.com/hexhoc/order-service/internal/entity"
	"github.com/hexhoc/order-service/internal/pb"
	"github.com/hexhoc/order-service/internal/repository"
	"github.com/hexhoc/order-service/pkg/kafka/publisher"
)

type OrderInterface interface {
	FindAll(ctx context.Context, request *pb.FindAllRequest) (*pb.FindAllResponse, error)
	FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error)
	Save(ctx context.Context, request *pb.SaveRequest) (*pb.StatusResponse, error)
	Update(ctx context.Context, request *pb.UpdateRequest) (*pb.StatusResponse, error)
	Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusResponse, error)
}

type OrderService struct {
	orderRepository       repository.OrderInterface
	paymentEventPublisher publisher.EventPublisher
	fetchGoodsPublisher   publisher.EventPublisher
	shipGoodsPublisher    publisher.EventPublisher
}

func NewOrderService(orderRepository repository.OrderInterface,
	paymentEventPublisher publisher.EventPublisher,
	fetchGoodsPublisher publisher.EventPublisher,
	shipGoodsPublisher publisher.EventPublisher) *OrderService {

	return &OrderService{
		orderRepository:       orderRepository,
		paymentEventPublisher: paymentEventPublisher,
		fetchGoodsPublisher:   fetchGoodsPublisher,
		shipGoodsPublisher:    shipGoodsPublisher,
	}
}

func (s *OrderService) FindAll(ctx context.Context, request *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	results, err := s.orderRepository.FindAll(ctx, request.WithItems, request.Limit, request.Offset)
	if err != nil {
		return &pb.FindAllResponse{
			Orders: []*pb.OrderDto{},
			Error:  err.Error(),
		}, err
	}

	var ordersList []*pb.OrderDto
	for _, order := range results {
		ordersList = append(ordersList, s.mapperToDto(order))
	}

	return &pb.FindAllResponse{
		Orders: ordersList,
		Error:  "",
	}, nil
}

func (s *OrderService) FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error) {
	result, err := s.orderRepository.FindById(ctx, request.WithItems, request.Id)
	if err != nil {
		return &pb.FindByIdResponse{
			Order: &pb.OrderDto{},
			Error: "",
		}, err
	}

	return &pb.FindByIdResponse{
		Order: s.mapperToDto(result),
		Error: "",
	}, nil
}

func (s *OrderService) Save(ctx context.Context, request *pb.SaveRequest) (*pb.StatusResponse, error) {
	id, err := s.orderRepository.Save(ctx, s.mapperToEntity(request.Order))
	if err != nil {
		return &pb.StatusResponse{Status: "NOT OK", Error: err.Error()}, err
	} else {

		//TODO: Add event publisher and send message to Retrieve paymant
		return &pb.StatusResponse{Status: id, Error: ""}, nil
	}
}

func (s *OrderService) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.StatusResponse, error) {
	err := s.orderRepository.Update(ctx, request.Id, s.mapperToEntity(request.Order))
	if err != nil {
		return &pb.StatusResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}
	return &pb.StatusResponse{
		Status: "OK",
		Error:  "",
	}, nil
}

func (s *OrderService) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusResponse, error) {
	err := s.orderRepository.Delete(ctx, request.Id)
	if err != nil {
		return &pb.StatusResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}
	return &pb.StatusResponse{
		Status: "OK",
		Error:  "",
	}, nil
}

func (s *OrderService) mapperToDto(e *entity.Order) *pb.OrderDto {
	dto := &pb.OrderDto{
		Id:         e.Id,
		CustomerId: e.CustomerId,
		Address:    e.Address,
		Status:     e.Status,
		IsDeleted:  e.IsDeleted,
		CreatedAt:  e.CreatedAt.Unix(),
		UpdatedAt:  e.UpdatedAt.Unix(),
		OrderItems: []*pb.OrderItemDto{},
	}

	for _, v := range e.OrderItems {
		price, _ := v.Price.Float64()
		dtoItem := &pb.OrderItemDto{
			Id:        v.Id,
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
			Price:     price,
		}
		dto.OrderItems = append(dto.OrderItems, dtoItem)
	}

	return dto
}

func (s *OrderService) mapperToEntity(d *pb.OrderDto) *entity.Order {
	e := &entity.Order{
		Id:         d.Id,
		CustomerId: d.CustomerId,
		Address:    d.Address,
		Status:     d.Status,
		IsDeleted:  d.IsDeleted,
		CreatedAt:  time.Unix(d.CreatedAt, 0),
		UpdatedAt:  time.Unix(d.UpdatedAt, 0),
		OrderItems: []*entity.OrderItem{},
	}

	for _, v := range d.OrderItems {
		item := &entity.OrderItem{
			Id:        v.Id,
			OrderId:   d.Id,
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
			Price:     *(big.NewFloat(v.Price)),
		}
		e.OrderItems = append(e.OrderItems, item)
	}

	return e
}
