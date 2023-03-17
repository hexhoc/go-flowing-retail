package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/product/pb"
	"google.golang.org/grpc"
)

type ProductService struct {
	Client pb.ProductServiceClient
}

func NewProductService(cc *grpc.ClientConn) *ProductService {
	productServiceClient := pb.NewProductServiceClient(cc)
	serviceClient := &ProductService{Client: productServiceClient}

	return serviceClient
}

func (svc *ProductService) FindAll(ctx *gin.Context) {
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

func (svc *ProductService) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	productRequest := &pb.FindByIdRequest{Id: uint32(id)}
	response, err := svc.Client.FindById(context.Background(), productRequest)
	if err != nil {
		//TODO: Сделать единый формат сообщения об ошибке
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ProductService) Save(ctx *gin.Context) {
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

func (svc *ProductService) SaveAll(ctx *gin.Context) {
	var requestBody pb.SaveAllRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := svc.Client.SaveAll(context.Background(), &requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ProductService) Update(ctx *gin.Context) {
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

func (svc *ProductService) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))
	request := &pb.DeleteRequest{Id: uint32(id)}
	response, err := svc.Client.Delete(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
