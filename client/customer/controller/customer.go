package controller

import (
	"grpc/client/customer/requests"
	"grpc/client/customer/responses"
	"grpc/helpers"
	"grpc/pb"
	"net/http"
	"time"

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

	newCustomer, err := h.Client.CreateCustomer(c.Request.Context(), &pb.CustomerModel{
		Name:       body.Name,
		IdPersonal: body.IdPersonal,
		Age:        body.Age,
		Address:    body.Address,
		Tags:       body.Tags,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["CREATE_CUSTOMER_ERROR"],
			ErrorMessage: "CREATE_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := responses.CustomerModelResponse{
		ID:         newCustomer.Id,
		Name:       newCustomer.Name,
		IdPersonal: newCustomer.IdPersonal,
		Age:        newCustomer.Age,
		Address:    newCustomer.Address,
		Tags:       newCustomer.Tags,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
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
	customer, err := h.Client.GetCustomer(c.Request.Context(), &pb.GetCustomerRequest{
		IdPersonal: body.IdPersonal,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["FIND_CUSTOMER_ERROR"],
			ErrorMessage: "FIND_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := responses.CustomerModelResponse{
		ID:         customer.Id,
		Name:       customer.Name,
		IdPersonal: customer.IdPersonal,
		Age:        customer.Age,
		Address:    customer.Address,
		Tags:       customer.Tags,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  customer.UpdatedAt,
	}

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
	newCustomer, err := h.Client.UpdateCustomer(c.Request.Context(), &pb.CustomerModel{
		Id:      body.ID,
		Name:    body.Name,
		Age:     body.Age,
		Address: body.Address,
		Tags:    body.Tags,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["UPDATE_CUSTOMER_ERROR"],
			ErrorMessage: "UPDATE_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := responses.CustomerModelResponse{
		ID:         newCustomer.Id,
		Name:       newCustomer.Name,
		IdPersonal: newCustomer.IdPersonal,
		Age:        newCustomer.Age,
		Address:    newCustomer.Address,
		Tags:       newCustomer.Tags,
		CreatedAt:  newCustomer.CreatedAt,
		UpdatedAt:  newCustomer.UpdatedAt,
	}

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

	rs, err := h.Client.DeleteCustomer(c.Request.Context(), &pb.CustomerModel{
		Id: body.ID,
	})
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
		customer := responses.CustomerModelResponse{
			ID:         item.Id,
			Name:       item.Name,
			IdPersonal: item.IdPersonal,
			Age:        item.Age,
			Tags:       item.Tags,
			Address:    item.Address,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		result = append(result, customer)
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
	customer, err := h.Client.AddTagsCustomer(c.Request.Context(), &pb.AddTagsCustomerRequest{
		Id:   body.ID,
		Tags: body.Tags,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["ADD_TAGS_CUSTOMER_ERROR"],
			ErrorMessage: "ADD_TAGS_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}

	result := responses.CustomerModelResponse{
		ID:        customer.Id,
		Name:      customer.Name,
		Age:       customer.Age,
		Address:   customer.Address,
		Tags:      customer.Tags,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
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
	customer, err := h.Client.DeleteTagsOfCustomer(c.Request.Context(), &pb.DeleteTagsOfCustomerRequest{
		Id:   body.ID,
		Tags: body.Tags,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.ErrorResponse{
			ErrorCode:    helpers.ErrorsCustomer["DELETE_TAGS_CUSTOMER_ERROR"],
			ErrorMessage: "DELETE_TAGS_CUSTOMER_ERROR",
			ErrorSystem:  err.Error(),
		})
		return
	}
	result := responses.CustomerModelResponse{
		ID:        customer.Id,
		Name:      customer.Name,
		Age:       customer.Age,
		Address:   customer.Address,
		Tags:      customer.Tags,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
	c.JSON(http.StatusOK, helpers.ResponseJSON{
		Status:  true,
		Message: "DELETE SUCCESS",
		Error:   "",
		Payload: result,
	})
	return
}
