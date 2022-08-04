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
	//Password string `json:"password"`
}

func Forgot(ctx *gin.Context, c pb.AuthServiceClient) {
	body := ForgotRequestBody{}
	ctx.Bind(&body)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//message := "You will receive a reset email if user with that email exist"

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email: body.Email,
		//Password: body.Password,

		/*
			res, err := c.Forgot(context.Background(), &pb.ForgotRequest{
				Email: body.Email,

		*/

	})

	if err != nil {
		//ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
