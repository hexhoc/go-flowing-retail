package app

import (
	"fmt"
	"net"

	"github.com/hexhoc/order-service/config"
	"github.com/hexhoc/order-service/internal/repository"
	"github.com/hexhoc/order-service/internal/service"
	"github.com/hexhoc/order-service/pkg/datasource/postgres"
	"github.com/hexhoc/order-service/pkg/logger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	logger.Init(cfg.LogLevel)
	log.Info("Starting order-service")

	pg, err := postgres.NewPostgresConnection(cfg.DBUrl, postgres.MaxPoolSize(1))
	if err != nil {
		log.Error(fmt.Errorf("app - Run - postgres.NewProductService: %w", err))
	}
	defer pg.Close()

	orderRepository := repository.NewOrderRepository(pg)
	orderItemRepository := repository.NewOrderItemRepository(pg)
	orderService := service.NewOrderService(orderRepository)
	orderItemService := service.NewOrderItemService(orderItemRepository)

	grpcServer := grpc.NewServer()
	// pb.RegisterProductServiceServer(grpcServer, productService)
	// pb.RegisterProductImageServiceServer(grpcServer, productImageService)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	log.Info("Order service start on ", cfg.Port)

	fmt.Println(orderService)
	fmt.Println(orderItemService)

}
