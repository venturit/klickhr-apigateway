package hris

import (
	"github.com/gin-gonic/gin"

	"klickhr-apigateway/pkg/config"
	"klickhr-apigateway/pkg/hris/routes"
)

//Initialize Routes and register
func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/hris")
	routes.POST("/uploadHRIS", svc.uploadHRIS)

	return svc
}

func (svc *ServiceClient) uploadHRIS(ctx *gin.Context) {
	routes.UploadHRIS(ctx, svc.Client)
}
