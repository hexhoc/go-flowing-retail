package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/auth"
	"github.com/hexhoc/api-gateway/internal/order"
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

	orderSvcConnect, err := grpc.Dial(cfg.OrderServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to orderServiceClient: ", err)
	}
	authSvc := auth.NewAuthService(authSvcConnect)
	middleware := auth.NewAuthMiddleware(authSvc)
	productSvc := product.NewProductService(productSvcConnect)
	imageSvc := product.NewImageService(productSvcConnect)
	orderSvc := order.NewOrderService(orderSvcConnect)

	auth.RegisterRoutes(r, authSvc)
	product.RegisterRoutes(r, middleware, productSvc, imageSvc)
	order.RegisterRoutes(r, middleware, orderSvc)

	r.Run(cfg.Port)
}
