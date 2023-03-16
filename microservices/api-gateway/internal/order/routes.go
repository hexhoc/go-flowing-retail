package order

import (
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth"
)

func RegisterRoutes(r *gin.Engine, middleware *auth.AuthMiddlewareConfig, orderSvc *OrderService) {
	orderGroup := r.Group("/api/v1/order")
	orderGroup.Use(middleware.AuthRequired)
	orderGroup.GET("/", orderSvc.FindAll)
	orderGroup.GET("/:id", orderSvc.FindById)
	orderGroup.POST("/", orderSvc.Save)
	orderGroup.PUT("/", orderSvc.Update)
	orderGroup.DELETE("/:id", orderSvc.Delete)
}
