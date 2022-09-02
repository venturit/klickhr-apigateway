package routes

import (
	"context"
	"net/http"

	"klickhr-apigateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

//Endpoint Recovery Password route
type RecoveryRequestBody struct {
	Password        string `json:"password"`
	Passwordconfirm string `json:"passwordconfirm"`
	OTP             string `json:"otp"`
}

type OtpRequestBody struct {
	OTP string `json:"otp"`
}

func ValidationOTP(ctx *gin.Context, c pb.AuthServiceClient) {
	body := OtpRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ValidateOTP(context.Background(), &pb.ValidationOTPRequest{
		Otp: body.OTP,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}

func Recovery(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RecoveryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Recovery(context.Background(), &pb.RecoveryRequest{
		Otp:             body.OTP,
		Password:        body.Password,
		Passwordconfirm: body.Passwordconfirm,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
