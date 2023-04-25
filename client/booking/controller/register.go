package controller

import (
	"grpc/pb"

	"github.com/gin-gonic/gin"
)

type BookingAPI interface {
	CreateTicket(c *gin.Context)
	FindTicket(c *gin.Context)
}

type BookingHandler struct {
	Client pb.BookingClient
}

func NewBookingAPI(client pb.BookingClient) BookingHandler {
	return BookingHandler{
		Client: client,
	}
}
