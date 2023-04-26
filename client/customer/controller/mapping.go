package controller

import (
	"grpc/client/customer/requests"
	"grpc/client/customer/responses"
	"grpc/pb"
)

func MappingCustomerResponse(customer *pb.CustomerModel) responses.CustomerModelResponse {
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
	return result
}

func MappingCreateCustomerRequest(body *requests.CustomerModelRequest) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Name:       body.Name,
		IdPersonal: body.IdPersonal,
		Age:        body.Age,
		Address:    body.Address,
		Tags:       body.Tags,
	}
	return result
}

func MappingUpdateCustomerRequest(body *requests.UpdateCustomerModelRequest) *pb.CustomerModel {
	result := &pb.CustomerModel{
		Id:      body.ID,
		Name:    body.Name,
		Age:     body.Age,
		Address: body.Address,
		Tags:    body.Tags,
	}
	return result
}

func MappingAddTagsCustomerRequest(body *requests.AddTagsCustomerRequest) *pb.AddTagsCustomerRequest {
	result := &pb.AddTagsCustomerRequest{
		Id:   body.ID,
		Tags: body.Tags,
	}
	return result
}

func MappingDeleteTagsCustomerRequest(body *requests.DeleteTagsOfCustomerRequest) *pb.DeleteTagsOfCustomerRequest {
	result := &pb.DeleteTagsOfCustomerRequest{
		Id:   body.ID,
		Tags: body.Tags,
	}
	return result
}
