package main

import (
	"log"

	"github.com/NicolasArizaR/klickhr-apigateway/pkg/auth"
	"github.com/NicolasArizaR/klickhr-apigateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	//authSvc := *auth.RegisterRoutes(r, &c)
	//product.RegisterRoutes(r, &c, &authSvc)
	//order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
