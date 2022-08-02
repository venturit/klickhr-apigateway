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
	routes.POST("/forgot", svc.Forgot)
	routes.PATCH("/recovery", svc.Recovery)
	//routes.GET("/refresh", svc.Recovery)
	//routes.GET("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	//routes.POST("/forgotpassword", rc.authController.ForgotPassword)
	//routes.PATCH("/resetpassword/:resetToken", rc.authController.ResetPassword)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Recovery(ctx *gin.Context) {
	routes.Recovery(ctx, svc.Client)
}

func (svc *ServiceClient) Forgot(ctx *gin.Context) {
	routes.Recovery(ctx, svc.Client)
}
