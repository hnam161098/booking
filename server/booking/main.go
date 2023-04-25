package main

import (
	"fmt"
	"grpc/config"
	"grpc/pb"
	"grpc/server/booking/controller"
	"grpc/server/booking/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	customerConnect, err := grpc.Dial(config.ADDRESS_SERVER["CUSTOMER_PORT"], grpc.WithInsecure())
	customerClient := pb.NewCustomerClient(customerConnect)

	bookingProcessData := handler.NewDBmanager()
	listen, err := net.Listen("tcp", config.ADDRESS_SERVER["BOOKING_PORT"])
	if err != nil {
		log.Fatal(err)
	}
	srv, err := controller.NewBooking(bookingProcessData, customerClient)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterBookingServer(server, srv)
	reflection.Register(server)
	fmt.Println("Booking service run on localhost:10002")
	server.Serve(listen)

}
