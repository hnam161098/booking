package controller

import (
	"fmt"
	"grpc/client/customer/requests"
	"grpc/client/customer/responses"
	"grpc/helpers"
	"grpc/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create customer
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	body := new(requests.CustomerModelRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}

	newCustomer, err := h.Client.CreateCustomer(c.Request.Context(), MappingCreateCustomerRequest(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["CREATE_CUSTOMER_ERROR"],
			ErrorMessage: "CREATE_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := MappingCustomerResponse(newCustomer)
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}

// get customer
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	body := new(requests.CustomerIdRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}
	customer, err := h.Client.GetCustomer(c.Request.Context(), &pb.GetCustomerRequest{IdPersonal: body.IdPersonal})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["FIND_CUSTOMER_ERROR"],
			ErrorMessage: "FIND_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := MappingCustomerResponse(customer)
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}

// update customer
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	body := new(requests.UpdateCustomerModelRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}
	customer, err := h.Client.UpdateCustomer(c.Request.Context(), MappingUpdateCustomerRequest(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["UPDATE_CUSTOMER_ERROR"],
			ErrorMessage: "UPDATE_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := MappingCustomerResponse(customer)
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}

// delete customer
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	body := new(requests.DeleteCustomerRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}

	rs, err := h.Client.DeleteCustomer(c.Request.Context(), &pb.CustomerModel{Id: body.ID})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["DELETE_CUSTOMER_ERROR"],
			ErrorMessage: "DELETE_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	if rs.Count == 0 {
		c.JSON(http.StatusOK, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["NOT_FOUND"],
			ErrorMessage: "NOT_FOUND",
			ErrorSystem:  "customer not found",
		})
		return
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "DELETE SUCCESS",
		Error:   "",
	})
	return
}

// get all customer
func (h *CustomerHandler) GetAllCustomer(c *gin.Context) {
	customers, err := h.Client.GetAllCustomer(c.Request.Context(), &pb.GetAllCustomerRequest{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["NOT_FOUND"],
			ErrorMessage: "NOT_FOUND",
			ErrorSystem:  err.Error(),
		})
		return
	}
	result := []responses.CustomerModelResponse{}
	for _, item := range customers.Customers {
		result = append(result, MappingCustomerResponse(item))
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}

// add tag customer
func (h *CustomerHandler) AddTagsCustomer(c *gin.Context) {
	body := new(requests.AddTagsCustomerRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}
	customer, err := h.Client.AddTagsCustomer(c.Request.Context(), MappingAddTagsCustomerRequest(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["ADD_TAGS_CUSTOMER_ERROR"],
			ErrorMessage: "ADD_TAGS_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := MappingCustomerResponse(customer)
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "SUCCESS",
		Error:   "",
		Payload: result,
	})
	return

}

// delete tag customer
func (h *CustomerHandler) DeleteTagsOfCustomer(c *gin.Context) {
	body := new(requests.DeleteTagsOfCustomerRequest)
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.Errors["MISSING_PARAMS"],
			ErrorMessage: "MISSING_PARAMS",
			ErrorSystem:  err.Error(),
		})
		return
	}
	customer, err := h.Client.DeleteTagsOfCustomer(c.Request.Context(), MappingDeleteTagsCustomerRequest(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["DELETE_TAGS_CUSTOMER_ERROR"],
			ErrorMessage: "DELETE_TAGS_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}
	fmt.Println("customer: ", customer)
	result := MappingCustomerResponse(customer)
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "DELETE SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}
