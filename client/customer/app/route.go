package app

import (
	"fmt"
	"grpc/config"
)

func Run() {
	// Register controller
	controller, g := RegisterCustomerClient()

	// API
	customer := g.Group("/customer")
	customer.POST("/create_customer", controller.CreateCustomer)
	customer.POST("/get_customer", controller.GetCustomer)
	customer.POST("/update_customer", controller.UpdateCustomer)
	customer.POST("/delete_customer", controller.DeleteCustomer)
	customer.POST("/delete_tags_customer", controller.DeleteTagsOfCustomer)
	customer.POST("/add_tags_customer", controller.AddTagsCustomer)
	customer.GET("/get_all_customer", controller.GetAllCustomer)

	fmt.Println("Customer API run port: 10001")
	// Run service
	g.Run(config.ADDRESS_CLIENT["CUSTOMER_PORT"])

}
