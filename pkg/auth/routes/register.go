package routes

import (
	"context"
	"net/http"

	"klickhr-apigateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

//Endpoint Register route cambiar struct por los necesarios al crear una account
type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	//EmployeeId int    `json:"employee_id"`
	//RoleType   int    `json:"role_type"`
	//Status bool `json:"status"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		//EmployeeId: int64(body.EmployeeId)
		//RoleType: int64(body.RoleType)
		//Status: body.Status,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
