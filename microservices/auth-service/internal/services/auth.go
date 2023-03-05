package services

import (
	"context"
	"github.com/hexhoc/auth-service/internal/pb"
)

type Server struct {
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	return &pb.RegisterResponse{
		Status: "ok",
		Error:  "",
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Status: "ok",
		Error:  "",
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {

	return &pb.ValidateResponse{
		Status: 0,
		Error:  "",
		UserId: 0,
	}, nil
}
