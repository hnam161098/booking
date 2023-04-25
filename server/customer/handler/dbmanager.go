package handler

import (
	"grpc/database_connection"

	"go.mongodb.org/mongo-driver/mongo"
)

type DbManager struct {
	collection *mongo.Collection
}

func NewDBmanager() CustomerHandler {
	var coll = database_connection.MongoDB.Db.Collection("customer")
	return &DbManager{
		collection: coll,
	}
}
