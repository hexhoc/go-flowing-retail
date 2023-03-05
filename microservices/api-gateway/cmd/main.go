package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth"
	"log"

	"github.com/hexhoc/api-gateway/config"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, c)

	fmt.Println("Starting api gateways")
	fmt.Println(authSvc)

	r.Run(c.Port)
}
