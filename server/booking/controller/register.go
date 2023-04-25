package controller

import (
	"grpc/pb"
	"grpc/server/booking/handler"
)

type BookingHandlerStruct struct {
	pb.UnimplementedBookingServer
	CustomerClient pb.CustomerClient
	BookingHandler handler.BookingHandler
}

func NewBooking(bookingHandler handler.BookingHandler, customerClient pb.CustomerClient) (*BookingHandlerStruct, error) {
	return &BookingHandlerStruct{
		CustomerClient: customerClient,
		BookingHandler: bookingHandler,
	}, nil
}
