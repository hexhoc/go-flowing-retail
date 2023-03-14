package product

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth"
)

func RegisterRoutes(r *gin.Engine, middleware *auth.AuthMiddlewareConfig, productSvc *ProductService, imageSvc *ImageService) {
	productGroup := r.Group("/api/v1/product")
	productGroup.Use(middleware.AuthRequired)
	productGroup.GET("/", productSvc.FindAll)
	productGroup.GET("/:id", productSvc.FindById)
	productGroup.POST("/", productSvc.Save)
	productGroup.POST("/batch", productSvc.SaveAll)
	productGroup.PUT("/", productSvc.Update)
	productGroup.DELETE("/:id", productSvc.Delete)

	//TODO: WARNING TEST
	imageGroup := r.Group("/api/v1/image")
	imageGroup.Use(middleware.AuthRequired)
	imageGroup.GET("/:productId", imageSvc.GetAllByProductId)
	imageGroup.POST("/:productId/upload", imageSvc.UploadImageToProduct)
	imageGroup.DELETE("/:productId/:imageName", imageSvc.GetAllByProductId)

}
