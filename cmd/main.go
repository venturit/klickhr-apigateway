package main

import (
	"log"

	"klickhr-apigateway/pkg/auth"
	"klickhr-apigateway/pkg/config"
	"klickhr-apigateway/pkg/hris"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	hris.RegisterRoutes(r, &c)
	//authSvc := *auth.RegisterRoutes(r, &c)
	//hris.RegisterRoutes(r, &c, &authSvc)
	//sessions.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
