package product

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/product/pb"
	"github.com/hexhoc/api-gateway/internal/product/routes"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func NewServiceClient(cfg *config.Config) *ServiceClient {
	cc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to productServiceClient: ", err)
	}

	productServiceClient := pb.NewProductServiceClient(cc)
	serviceClient := &ServiceClient{Client: productServiceClient}

	return serviceClient
}

func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	routes.FindAll(ctx, svc.Client, uint32(limit), uint32(offset))
}

func (svc *ServiceClient) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	routes.FindById(ctx, svc.Client, uint32(id))
}
