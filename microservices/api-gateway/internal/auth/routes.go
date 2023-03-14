package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, svc *AuthService) {
	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
}
