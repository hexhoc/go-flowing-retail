package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
	"google.golang.org/grpc"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	cc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to productServiceClient: ", err)
	}
	productSvc := NewProductService(cc)
	imageSvc := NewImageService(cc)

	productGroup := r.Group("/api/v1/product")
	productGroup.GET("/", productSvc.FindAll)
	productGroup.GET("/:id", productSvc.FindById)
	productGroup.POST("/", productSvc.Save)
	productGroup.POST("/batch", productSvc.SaveAll)
	productGroup.PUT("/", productSvc.Update)
	productGroup.DELETE("/:id", productSvc.Delete)

	imageGroup := r.Group("/api/v1/image")
	imageGroup.GET("/:productId", imageSvc.GetAllByProductId)
	imageGroup.POST("/:productId/upload", imageSvc.UploadImageToProduct)
	imageGroup.DELETE("/:productId/:imageName", imageSvc.GetAllByProductId)

}
