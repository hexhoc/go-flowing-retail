package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/auth"
	"github.com/hexhoc/api-gateway/internal/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, c)
	productSvc := product.RegisterRoutes(r, c)

	fmt.Println("Starting api gateways")
	fmt.Println(authSvc)
	fmt.Println(productSvc)

	r.Run(c.Port)
}
