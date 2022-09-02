package auth

import (
	"github.com/gin-gonic/gin"

	"klickhr-apigateway/pkg/auth/routes"

	"klickhr-apigateway/pkg/config"
)

//Initialize Routes and register
func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/forgotpassword", svc.Forgot)
	routes.POST("/validationOTP", svc.ValidationOTP)
	routes.PATCH("/recoverypassword", svc.Recovery)
	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Forgot(ctx *gin.Context) {
	routes.Forgot(ctx, svc.Client)
}

func (svc *ServiceClient) ValidationOTP(ctx *gin.Context) {
	routes.ValidationOTP(ctx, svc.Client)
}

func (svc *ServiceClient) Recovery(ctx *gin.Context) {
	routes.Recovery(ctx, svc.Client)
}
