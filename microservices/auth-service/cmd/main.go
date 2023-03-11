package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hexhoc/auth-service/config"
	"github.com/hexhoc/auth-service/internal/db"
	"github.com/hexhoc/auth-service/internal/pb"
	"github.com/hexhoc/auth-service/internal/services"
	"github.com/hexhoc/auth-service/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting auth-service")

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "auth-service",
		ExpirationHours: 24 * 365,
	}

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	fmt.Println("Auth service start on ", c.Port)
}
