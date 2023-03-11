package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/product/pb"
)

func FindAll(ctx *gin.Context, productClient pb.ProductServiceClient, limit uint32, offset uint32) {
	productRequest := &pb.FindAllRequest{Limit: limit, Offset: offset}
	response, err := productClient.FindAll(context.Background(), productRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func FindById(ctx *gin.Context, productClient pb.ProductServiceClient, id uint32) {
	productRequest := &pb.FindByIdRequest{Id: id}
	response, err := productClient.FindById(context.Background(), productRequest)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
