package app

import (
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
	var productImageRepository repository.ProductImageInterface = repository.NewProductImageRepository(pg)
	var productService service.ProductInterface = service.NewProductService(productRepository)
	var productImageService service.ProductImageInterface = service.NewProductImageService(productImageRepository)

	fmt.Println(productService)
	fmt.Println(productImageService)

	//product := &entity.Product{
	//	Name:          "test",
	//	Intro:         "test",
	//	Description:   "test",
	//	CategoryId:    1,
	//	Category:      entity.Category{},
	//	OriginalPrice: 10,
	//	SellingPrice:  10,
	//	IsSale:        true,
	//	IsDeleted:     false,
	//	CreatedAt:     time.Now(),
	//	UpdatedAt:     time.Now(),
	//}
	//err = productService.Save(context.Background(), product)
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//products, err := productService.FindAll(context.Background())
	//if err != nil {
	//	log.Error(err)
	//}
	//for i := 0; i < len(products); i++ {
	//	fmt.Println(products[i])
	//}
	//
	//createProduct, err := productService.FindById(context.Background(), 8)
	//if err != nil {
	//	log.Error(err)
	//}
	//fmt.Println(createProduct)
	//
	//err = productService.Delete(context.Background(), 8)
	//if err != nil {
	//	log.Error(err)
	//}
}
