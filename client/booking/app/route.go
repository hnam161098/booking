package app

import (
	"fmt"
	"grpc/config"
)

func Run() {
	// Register controller
	controller, g := RegisterBookingClient()
	
	// API route
	booking := g.Group("/booking")
	booking.POST("/create-ticket", controller.CreateTicket)
	booking.POST("/find-ticket", controller.FindTicket)

	fmt.Println("Booking API run port: 20002")
	g.Run(config.ADDRESS_CLIENT["BOOKING_PORT"])
}
