package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth"
	"log"

	"github.com/hexhoc/api-gateway/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &config)

	fmt.Println("Starting api gateways")
	fmt.Println(authSvc)
}
