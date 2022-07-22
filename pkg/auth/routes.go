package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/NicolasArizaR/klickhr-apigateway/pkg/auth/routes"
	"github.com/NicolasArizaR/klickhr-apigateway/pkg/config"
)

//Initialize Routes and register
func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
