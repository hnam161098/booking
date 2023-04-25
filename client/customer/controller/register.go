package controller

import (
	"grpc/pb"

	"github.com/gin-gonic/gin"
)

// Register API
type CustomerAPI interface {
	CreateCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
	GetAllCustomer(c *gin.Context)
	AddTagsCustomer(c *gin.Context)
	DeleteTagsOfCustomer(c *gin.Context)
}

type CustomerHandler struct {
	Client pb.CustomerClient
}

func NewCustomerAPI(client pb.CustomerClient) CustomerHandler {
	return CustomerHandler{
		Client: client,
	}
}
