package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := NewServiceClient(c)

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}
