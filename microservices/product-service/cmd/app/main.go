package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hexhoc/product-service/config"
	"github.com/hexhoc/product-service/internal/repository"
	"github.com/hexhoc/product-service/internal/usecase"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
)

func main() {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	l.Println("Starting product-service")

	c, err := config.LoadConfig()
	if err != nil {
		l.Fatalln("Failed at config", err)
	}

	pg, err := postgres.NewPostgresConnection(c.DBUrl, postgres.MaxPoolSize(1))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	productRepository := repository.NewProductRepository(pg)
	productUseCase := usecase.NewProductUseCase(productRepository)

	products := productUseCase.FindAll()
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

}
