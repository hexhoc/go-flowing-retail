package product

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := NewServiceClient(c)

	routes := r.Group("/product")
	routes.GET("/", svc.FindAll)
	routes.GET("/:id", svc.FindById)
	routes.POST("/", svc.Save)
	routes.POST("/batch", svc.SaveAll)
	routes.PUT("/", svc.Update)
	routes.DELETE("/:id", svc.Delete)

	return svc
}
