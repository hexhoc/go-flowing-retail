package main

import (
	"fmt"
	"log"

	"github.com/hexhoc/product-service/config"
	"github.com/hexhoc/product-service/internal/interfaces"
	"github.com/hexhoc/product-service/internal/repository"
	"github.com/hexhoc/product-service/internal/usecase"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
)

func main() {
	fmt.Println("Starting product-service")

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	pg, err := postgres.NewPostgresConnection(c.DBUrl, postgres.MaxPoolSize(1))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	var productRepository interfaces.ProductRepository = repository.NewProductRepository(pg)
	var productUseCase interfaces.ProductUseCase = usecase.NewProductUseCase(productRepository)

	products := productUseCase.FindAll()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

}
