package app

import (
	"grpc/client/booking/controller"
	"grpc/config"
	"grpc/pb"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RegisterBookingClient() (controller.BookingHandler, *gin.Engine) {
	connect, err := grpc.Dial(config.ADDRESS_SERVER["BOOKING_PORT"], grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBookingClient(connect)

	b := controller.NewBookingAPI(client)
	g := gin.Default()
	return b, g
}
