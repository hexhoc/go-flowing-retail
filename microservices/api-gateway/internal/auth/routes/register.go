package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hexhoc/api-gateway/internal/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, authClient pb.AuthServiceClient) {
	request := RegisterRequestBody{}
	err := ctx.BindJSON(request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := authClient.Register(context.Background(), &pb.RegisterRequest{Email: request.Email, Password: request.Password})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
