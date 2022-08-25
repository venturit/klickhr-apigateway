package main

import (
	"log"

	"klickhr-apigateway/pkg/auth"
	"klickhr-apigateway/pkg/config"

	//"klickhr.srm.apigateway/pkg/employee"
	//"klickhr.srm.apigateway/pkg/hris"
	//"klickhr.srm.apigateway/pkg/organization"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)
	//hris.RegisterRoutes(r, &c)
	//organization.RegisterRoutes(r, &c)
	//employee.RegisterRoutes(r, &c)

	/*
			authSvc := *auth.RegisterRoutes(r, &c)
		    hris.RegisterRoutes(r, &c, &authSvc)
		    organization.RegisterRoutes(r, &c, &authSvc)
	*/

	r.Run(c.Port)
}
