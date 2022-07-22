package routes

import (
	"context"
	"net/http"

	"github.com/NicolasArizaR/klickhr-apigateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

//Endpoint Login Route
type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Function Login returna JWT
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
