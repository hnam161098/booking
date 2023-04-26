package controller

import (
	"context"
	"errors"
	"grpc/helpers"
	"grpc/pb"
	"grpc/server/customer/models"
	"log"
	"time"
)

func MappingResultEStoCustomerPbModel(m map[string]interface{}) *pb.CustomerModel {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("message panic: %v", r)
			return
		}
	}()

	rsES := m["_source"].(map[string]interface{})
	listTags := []string{}

	for _, item := range rsES["tags"].([]interface{}) {
		listTags = append(listTags, item.(string))
	}
	result := &pb.CustomerModel{
		Id:         rsES["id"].(string),
		Name:       rsES["name"].(string),
		IdPersonal: rsES["id_personal"].(string),
		Age:        rsES["age"].(string),
		Address:    rsES["address"].(string),
		Tags:       listTags,
		CreatedAt:  rsES["created_at"].(string),
		UpdatedAt:  rsES["updated_at"].(string),
	}
	return result

}

func (c *CustomerHandlerStruct) CreateCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.CustomerModel, error) {
	req := models.Customer{
		Name:       in.Name,
		IdPersonal: in.IdPersonal,
		Age:        in.Age,
		Address:    in.Address,
		Tags:       in.Tags,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	customer, err := c.CustomerHandler.CreateCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}

	// sync to ES
	go func() {
		reqES := models.CustomerESModel{
			ID:         customer.ID,
			Name:       customer.Name,
			IdPersonal: customer.IdPersonal,
			Age:        customer.Age,
			Address:    customer.Address,
			Tags:       customer.Tags,
			CreatedAt:  customer.CreatedAt,
			UpdatedAt:  customer.UpdatedAt,
		}
		err1 := helpers.InsertDocumentES(models.IndexCustomer, reqES, customer.ID)
		if err1 != nil {
			log.Println(err1)
		}
	}()

	result := MappingCreateCustomerToPbModel(customer)
	return result, nil
}

func (c *CustomerHandlerStruct) GetCustomer(ctx context.Context, in *pb.GetCustomerRequest) (*pb.CustomerModel, error) {
	// get from Elasticsearch
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"id_personal": map[string]interface{}{
					"value": in.IdPersonal,
				},
			},
		},
	}
	rsES, err := helpers.QueryES(models.IndexCustomer, query)
	if err != nil {
		return nil, err
	}
	if rsES["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64) == 0 {
		return nil, errors.New("Customer not exist")
	}
	customer := MappingResultEStoCustomerPbModel(rsES["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{}))
	return customer, nil
}

func (c *CustomerHandlerStruct) UpdateCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.CustomerModel, error) {

	req := models.Customer{
		ID:      in.Id,
		Name:    in.Name,
		Age:     in.Age,
		Address: in.Address,
		Tags:    in.Tags,
	}
	customer, err := c.CustomerHandler.UpdateCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	// sync ES
	go func() {
		queryES := map[string]interface{}{
			"id":          customer.ID,
			"name":        customer.Name,
			"id_personal": customer.IdPersonal,
			"age":         customer.Age,
			"address":     customer.Address,
			"tags":        customer.Tags,
			"created_at":  customer.CreatedAt,
			"updated_at":  customer.UpdatedAt,
		}
		errUpdate := helpers.UpdateES(models.IndexCustomer, queryES, customer.ID)
		if errUpdate != nil {
			log.Println(errUpdate)
		}
	}()

	result := MappingUpdateCustomerToPbModel(customer)
	return result, nil
}

func (c *CustomerHandlerStruct) DeleteCustomer(ctx context.Context, in *pb.CustomerModel) (*pb.DeleteCustomerResponse, error) {
	req := models.Customer{
		ID: in.Id,
	}
	// delete in mongoDB
	count, err := c.CustomerHandler.DeleteCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	// delete in ES
	go func() {
		queryES := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"match": map[string]interface{}{
							"id": in.Id,
						},
					},
				},
			},
		}
		errES := helpers.DeleteES(models.IndexCustomer, queryES)
		if errES != nil {
			log.Println(errES)
		}
	}()
	return &pb.DeleteCustomerResponse{
		Count: int64(count),
	}, nil
}

func (c *CustomerHandlerStruct) GetAllCustomer(ctx context.Context, in *pb.GetAllCustomerRequest) (*pb.AllCustomerResponse, error) {
	// get from ES
	queryES := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	totalDocs, err := helpers.QueryES(models.IndexCustomer, queryES)
	if err != nil {
		return nil, err
	}
	var result []*pb.CustomerModel
	listDataES := totalDocs["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, customer := range listDataES {
		listTags := []string{}
		if customer.(map[string]interface{})["_source"].(map[string]interface{})["tags"] != nil {
			for _, item := range customer.(map[string]interface{})["_source"].(map[string]interface{})["tags"].([]interface{}) {
				listTags = append(listTags, item.(string))
			}
		}
		cus := MappingResultEStoCustomerPbModel(customer.(map[string]interface{}))
		result = append(result, cus)
	}
	return &pb.AllCustomerResponse{
		Customers: result,
	}, nil
}

func (c *CustomerHandlerStruct) AddTagsCustomer(ctx context.Context, in *pb.AddTagsCustomerRequest) (*pb.CustomerModel, error) {
	req := models.AddTags{
		ID:   in.Id,
		Tags: in.Tags,
	}
	customer, err := c.CustomerHandler.AddTagsCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	// sync to ES
	go func() {
		queryES := map[string]interface{}{
			"id":         customer.ID,
			"tags":       customer.Tags,
			"updated_at": customer.UpdatedAt,
		}
		errUpdate := helpers.UpdateES(models.IndexCustomer, queryES, customer.ID)
		if errUpdate != nil {
			log.Println(errUpdate)
		}
	}()
	result := MappingAddTagsCustomerToPbModel(customer)
	return result, nil
}

func (c *CustomerHandlerStruct) DeleteTagsOfCustomer(ctx context.Context, in *pb.DeleteTagsOfCustomerRequest) (*pb.CustomerModel, error) {
	req := models.DeleteTags{
		ID:   in.Id,
		Tags: in.Tags,
	}
	customer, err := c.CustomerHandler.DeleteTagsOfCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	result := MappingDeleteTagsCustomerToPbModel(customer)
	return result, nil

}
