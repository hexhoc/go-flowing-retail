package product

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := NewServiceClient(c)

	routes := r.Group("/product")
	routes.GET("/", svc.FindAll)

	return svc
}
