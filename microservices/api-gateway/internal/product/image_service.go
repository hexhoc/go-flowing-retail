package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/product/pb"
	"google.golang.org/grpc"
)

type ImageService struct {
	Client pb.ProductImageServiceClient
}

func NewImageService(cc *grpc.ClientConn) *ImageService {
	ImageServiceClient := pb.NewProductImageServiceClient(cc)
	serviceClient := &ImageService{Client: ImageServiceClient}

	return serviceClient
}

func (svc *ImageService) GetAllByProductId(ctx *gin.Context) {
	productId, _ := strconv.Atoi(ctx.Param("productId"))

	request := &pb.GetAllRequest{
		Id: uint32(productId),
	}
	response, err := svc.Client.GetAllByProductId(context.Background(), request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ImageService) UploadImageToProduct(ctx *gin.Context) {
	productId, _ := strconv.Atoi(ctx.Param("productId"))
	file, _ := ctx.FormFile("file")
	r, _ := file.Open()
	fileBytes := make([]byte, file.Size)
	r.Read(fileBytes)

	request := &pb.UploadImageRequest{
		File:      fileBytes,
		Filename:  file.Filename,
		ProductId: uint32(productId),
	}
	response, err := svc.Client.UploadImageToProduct(context.Background(), request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (svc *ImageService) DeleteByNameAndProductId(ctx *gin.Context) {
	productId, _ := strconv.Atoi(ctx.Param("productId"))
	imageName := ctx.Param("imageName")
	request := &pb.DeleteImageRequest{
		ImageName: imageName,
		ProductId: uint32(productId),
	}

	response, err := svc.Client.DeleteByNameAndProductId(context.Background(), request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// rpc (DeleteImageRequest) returns (DeleteImageResponse) {}
