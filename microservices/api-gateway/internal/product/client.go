package product

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"github.com/hexhoc/api-gateway/internal/product/pb"
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

	productRequest := &pb.FindAllRequest{Limit: uint32(limit), Offset: uint32(offset)}
	response, err := svc.Client.FindAll(context.Background(), productRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ServiceClient) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	productRequest := &pb.FindByIdRequest{Id: uint32(id)}
	response, err := svc.Client.FindById(context.Background(), productRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ServiceClient) Save(ctx *gin.Context) {
	var requestBody pb.SaveRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := svc.Client.Save(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ServiceClient) SaveAll(ctx *gin.Context) {
	var requestBody pb.SaveAllRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := svc.Client.SaveAll(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	var requestBody pb.UpdateRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := svc.Client.Update(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	request := &pb.DeleteRequest{Id: uint32(id)}
	response, err := svc.Client.Delete(context.Background(), request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
