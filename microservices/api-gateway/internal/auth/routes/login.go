package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, authClient pb.AuthServiceClient) {
	request := LoginRequestBody{}
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := authClient.Login(context.Background(), &pb.LoginRequest{
		Email:    request.Email,
		Password: request.Password})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"status": false, "message": err.Error()})
	}

	ctx.JSON(http.StatusCreated, response)
}
