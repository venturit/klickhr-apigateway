package routes

import (
	"context"
	"net/http"

	"klickhr-apigateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

//Endpoint Forgot Password route
type ForgotRequestBody struct {
	Email string `json:"email"`
}

func Forgot(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ForgotRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//Receive email and send it to the forgotpassword route
	res, err := c.Forgot(context.Background(), &pb.ForgotRequest{
		Email: body.Email,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
