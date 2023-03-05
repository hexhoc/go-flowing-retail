package main

import (
	"fmt"
	"github.com/hexhoc/auth-service/config"
	"github.com/hexhoc/auth-service/internal/pb"
	"github.com/hexhoc/auth-service/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting auth-service")

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	s := services.Server{}
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
