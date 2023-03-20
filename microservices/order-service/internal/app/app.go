package app

import (
	"context"
	"fmt"
	"net"

	"github.com/hexhoc/order-service/config"
	"github.com/hexhoc/order-service/internal/pb"
	"github.com/hexhoc/order-service/internal/repository"
	"github.com/hexhoc/order-service/internal/service"
	"github.com/hexhoc/order-service/pkg/datasource/postgres"
	"github.com/hexhoc/order-service/pkg/kafka/consumer"
	"github.com/hexhoc/order-service/pkg/kafka/publisher"
	"github.com/hexhoc/order-service/pkg/logger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	// LOGGER INIT
	logger.Init(cfg.LogLevel)
	log.Info("Starting order-service")

	// POSTGRES INIT
	pg, err := postgres.NewPostgresConnection(cfg.DBUrl, postgres.MaxPoolSize(1))
	if err != nil {
		log.Error(fmt.Errorf("app - Run - postgres.NewProductService: %w", err))
	}
	defer pg.Close()

	// KAFKA PUBLISHER INIT
	paymentEventPublisher := publisher.NewPublisher([]string{cfg.KafkaAddr}, "paymentTopic")
	fetchGoodsPublisher := publisher.NewPublisher([]string{cfg.KafkaAddr}, "fetchGoodsTopic")
	shipGoodsPublisher := publisher.NewPublisher([]string{cfg.KafkaAddr}, "shipGoodsTopic")

	// USECASE AND REPOSITORY INIT
	orderRepository := repository.NewOrderRepository(pg)
	orderService := service.NewOrderService(orderRepository, paymentEventPublisher, fetchGoodsPublisher, shipGoodsPublisher)

	// KAFKA CONSUMER
	consumer.Consume(context.Background(), []string{cfg.KafkaAddr}, "paymentTopic", orderService.EventHandle)
	consumer.Consume(context.Background(), []string{cfg.KafkaAddr}, "fetchGoodsTopic", orderService.EventHandle)
	consumer.Consume(context.Background(), []string{cfg.KafkaAddr}, "shipGoodsTopic", orderService.EventHandle)

	// GRPC SERVER
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	log.Info("Order service start on ", cfg.Port)
}
