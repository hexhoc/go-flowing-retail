package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/auth"
	"github.com/hexhoc/api-gateway/internal/product"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvcConnect, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to authServiceClient: ", err)
	}

	productSvcConnect, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to productServiceClient: ", err)
	}
	authSvc := auth.NewAuthService(authSvcConnect)
	middleware := auth.NewAuthMiddleware(authSvc)
	productSvc := product.NewProductService(productSvcConnect)
	imageSvc := product.NewImageService(productSvcConnect)

	auth.RegisterRoutes(r, authSvc)
	product.RegisterRoutes(r, middleware, productSvc, imageSvc)

	fmt.Println("Starting api gateways")

	r.Run(cfg.Port)
}
