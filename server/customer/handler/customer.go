package handler

import (
	"context"
	"grpc/server/customer/models"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type CustomerHandler interface {
	CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	GetCustomer(ctx context.Context, model *models.CustomerRequest) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, model *models.Customer) (int, error)
	GetAllCustomer(ctx context.Context) ([]*models.Customer, error)
	AddTagsCustomer(ctx context.Context, model *models.AddTags) (*models.Customer, error)
	DeleteTagsOfCustomer(ctx context.Context, model *models.DeleteTags) (*models.Customer, error)
}

func (m *DbManager) CreateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {
	filter := bson.M{
		"name":        model.Name,
		"id_personal": model.IdPersonal,
		"age":         model.Age,
		"address":     model.Address,
		"tags":        model.Tags,
		"created_at":  model.CreatedAt,
		"updated_at":  model.UpdatedAt,
	}

	// check exist
	filterCheck := bson.M{
		"id_personal": model.IdPersonal,
	}
	var customerExist models.Customer
	errF := m.collection.FindOne(context.TODO(), filterCheck).Decode(&customerExist)
	if errF == nil || customerExist.Name != "" {
		return &models.Customer{}, errors.New("Customer aready exist")
	}

	rs, err := m.collection.InsertOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:         rs.InsertedID.(primitive.ObjectID).Hex(),
		Name:       model.Name,
		IdPersonal: model.IdPersonal,
		Age:        model.Age,
		Address:    model.Address,
		Tags:       model.Tags,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}, nil
}

func (m *DbManager) GetCustomer(ctx context.Context, model *models.CustomerRequest) (*models.Customer, error) {
	var result models.Customer

	filter := bson.M{
		"id_personal": model.IdPersonal,
	}
	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil

}

func CheckDuplicate(sli []string, s string) bool {
	for _, v := range sli {
		if v == s {
			return true
		}
	}
	return false
}

func (m *DbManager) UpdateCustomer(ctx context.Context, model *models.Customer) (*models.Customer, error) {

	id, _ := primitive.ObjectIDFromHex(model.ID)

	filter := bson.M{
		"_id": id,
	}

	var cus models.Customer
	_ = m.collection.FindOne(context.TODO(), filter).Decode(&cus)

	if model.Name == "" {
		model.Name = cus.Name
	}
	if model.Age == "" {
		model.Age = cus.Age
	}
	if model.Address == "" {
		model.Address = cus.Address
	}

	update := bson.M{
		"$set": bson.M{
			"name":       model.Name,
			"age":        model.Age,
			"address":    model.Address,
			"updated_at": time.Now().Format(time.RFC3339),
		},
	}

	var customer *models.Customer
	opt := options.FindOneAndUpdateOptions{}
	opt.SetReturnDocument(1)

	err := m.collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return customer, nil

}

func (m *DbManager) DeleteCustomer(ctx context.Context, model *models.Customer) (int, error) {
	id, _ := primitive.ObjectIDFromHex(model.ID)

	filter := bson.M{
		"_id": id,
	}

	rs, err := m.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return int(rs.DeletedCount), nil

}

func (m *DbManager) GetAllCustomer(ctx context.Context) ([]*models.Customer, error) {
	var customers []*models.Customer

	cursor, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &customers); err != nil {
		return nil, err
	}
	return customers, nil

}

func (m *DbManager) AddTagsCustomer(ctx context.Context, model *models.AddTags) (*models.Customer, error) {
	id, _ := primitive.ObjectIDFromHex(model.ID)
	filter := bson.M{
		"_id": id,
	}

	var cus models.Customer
	_ = m.collection.FindOne(context.TODO(), filter).Decode(&cus)
	if len(model.Tags) == 0 {
		model.Tags = cus.Tags
	} else {
		for _, v := range model.Tags {
			if CheckDuplicate(cus.Tags, v) == false {
				cus.Tags = append(cus.Tags, v)
			}
		}
		model.Tags = cus.Tags
	}
	updateQuery := bson.M{
		"$set": bson.M{
			"tags":       model.Tags,
			"updated_at": time.Now().Format(time.RFC3339),
		},
	}

	opt := options.FindOneAndUpdateOptions{}
	opt.SetReturnDocument(1)

	var customer models.Customer
	err := m.collection.FindOneAndUpdate(ctx, filter, updateQuery, &opt).Decode(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (m *DbManager) DeleteTagsOfCustomer(ctx context.Context, model *models.DeleteTags) (*models.Customer, error) {
	var customer models.Customer
	id, _ := primitive.ObjectIDFromHex(model.ID)
	filter := bson.M{
		"_id": id,
	}

	deleteQuery := bson.M{
		"$pull": bson.M{
			"tags": bson.M{
				"$in": model.Tags,
			},
		},
	}
	_, err := m.collection.UpdateOne(ctx, filter, deleteQuery)
	if err != nil {
		return nil, err
	}

	err1 := m.collection.FindOne(ctx, filter).Decode(&customer)
	if err1 != nil {
		return nil, err1
	}

	return &customer, nil

}
