package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth/pb"
	"github.com/hexhoc/api-gateway/internal/auth/routes"
	"google.golang.org/grpc"
)

type AuthService struct {
	Client pb.AuthServiceClient
}

func NewAuthService(cc *grpc.ClientConn) *AuthService {
	authServiceClient := pb.NewAuthServiceClient(cc)
	serviceClient := &AuthService{Client: authServiceClient}

	return serviceClient
}

func (svc *AuthService) Login(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *AuthService) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
