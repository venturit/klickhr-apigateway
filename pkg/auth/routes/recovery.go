package routes

import (
	"context"
	"net/http"

	"klickhr-apigateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

//Endpoint Recovery Password route
type RecoveryRequestBody struct {
	Email string `json:"email"`
	//otp   int64  `json:"password"`
}

func Recovery(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RecoveryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email: body.Email,
		//otp:   body.otp,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
