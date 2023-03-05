package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/auth/pb"
	"github.com/hexhoc/api-gateway/internal/auth/routes"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func NewServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.AuthServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to authServiceClient: ", err)
	}

	authServiceClient := pb.NewAuthServiceClient(cc)
	serviceClient := &ServiceClient{Client: authServiceClient}

	return serviceClient
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}
