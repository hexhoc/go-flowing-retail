package app

import (
	"fmt"
	"net"

	"github.com/hexhoc/product-service/config"
	"github.com/hexhoc/product-service/internal/pb"
	"github.com/hexhoc/product-service/internal/repository"
	"github.com/hexhoc/product-service/internal/service"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
	"github.com/hexhoc/product-service/pkg/logger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) {
	logger.Init(cfg.LogLevel)
	log.Info("Starting product-service")

	pg, err := postgres.NewPostgresConnection(cfg.DBUrl, postgres.MaxPoolSize(1))
	if err != nil {
		log.Error(fmt.Errorf("app - Run - postgres.NewProductService: %w", err))
	}
	defer pg.Close()

	var productRepository repository.ProductInterface = repository.NewProductRepository(pg)
	var productImageRepository repository.ProductImageInterface = repository.NewProductImageRepository(pg)
	var productService service.ProductInterface = service.NewProductService(productRepository)
	var productImageService service.ProductImageInterface = service.NewProductImageService(productImageRepository)

	fmt.Println(productImageService)

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)
	pb.RegisterProductImageServiceServer(grpcServer, productImageService)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	log.Info("Product service start on ", cfg.Port)

}
