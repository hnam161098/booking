package app

import (
	"grpc/client/customer/controller"
	"grpc/config"
	"grpc/pb"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RegisterCustomerClient() (controller.CustomerHandler, *gin.Engine) {
	connCustomerServer, err := grpc.Dial(config.ADDRESS_SERVER["CUSTOMER_PORT"], grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	connectCustomerClient := pb.NewCustomerClient(connCustomerServer)

	c := controller.NewCustomerAPI(connectCustomerClient)
	g := gin.Default()
	return c, g
}
