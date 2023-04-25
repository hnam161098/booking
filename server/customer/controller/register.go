package controller

import (
	"grpc/pb"
	"grpc/server/customer/handler"
)

type CustomerHandlerStruct struct {
	pb.UnimplementedCustomerServer
	CustomerHandler handler.CustomerHandler
}

func NewCustomer(customerHandle handler.CustomerHandler) (*CustomerHandlerStruct, error) {
	return &CustomerHandlerStruct{
		CustomerHandler: customerHandle,
	}, nil
}
