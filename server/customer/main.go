package main

import (
	"fmt"
	"grpc/config"
	"grpc/pb"
	"grpc/server/customer/controller"
	"grpc/server/customer/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", config.ADDRESS_SERVER["CUSTOMER_PORT"])
	if err != nil {
		log.Fatal(err)
	}
	customerProcessData := handler.NewDBmanager()

	srv, err := controller.NewCustomer(customerProcessData)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterCustomerServer(server, srv)
	reflection.Register(server)
	fmt.Println("Customer service run on localhost:10001")
	server.Serve(listen)

}
