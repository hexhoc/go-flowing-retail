package order

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/order/pb"
	"google.golang.org/grpc"
)

type OrderService struct {
	Client pb.OrderServiceClient
}

func NewOrderService(cc *grpc.ClientConn) *OrderService {
	orderServiceClient := pb.NewOrderServiceClient(cc)
	serviceClient := &OrderService{Client: orderServiceClient}

	return serviceClient
}

func (svc *OrderService) FindAll(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	productRequest := &pb.FindAllRequest{Limit: uint32(limit), Offset: uint32(offset)}
	response, err := svc.Client.FindAll(context.Background(), productRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *OrderService) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	request := &pb.FindByIdRequest{Id: id}
	response, err := svc.Client.FindById(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *OrderService) Save(ctx *gin.Context) {
	var requestBody pb.SaveRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := svc.Client.Save(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *OrderService) Update(ctx *gin.Context) {
	var requestBody pb.UpdateRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := svc.Client.Update(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *OrderService) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	request := &pb.DeleteRequest{Id: id}
	response, err := svc.Client.Delete(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
