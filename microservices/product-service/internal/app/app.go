package app

import (
	"context"
	"fmt"
	"github.com/hexhoc/product-service/config"
	"github.com/hexhoc/product-service/internal/repository"
	"github.com/hexhoc/product-service/internal/service"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
	"github.com/hexhoc/product-service/pkg/logger"
	log "github.com/sirupsen/logrus"
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
	var productService service.ProductInterface = service.NewProductService(productRepository)

	products, err := productService.FindAll(context.Background())
	if err != nil {
		log.Error(err)
	}
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}
}
