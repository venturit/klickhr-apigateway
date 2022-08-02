package hris

import (
	"fmt"

	"klickhr-apigateway/pkg/config"
	"klickhr-apigateway/pkg/hris/pb"

	"google.golang.org/grpc"
)

//Hris Microservice Client
type ServiceClient struct {
	Client pb.HRISServiceClient
}

func InitServiceClient(c *config.Config) pb.HRISServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.HRISSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewHRISServiceClient(cc)
}
