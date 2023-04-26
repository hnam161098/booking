package database_connection

import (
	"context"
	"fmt"
	"grpc/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoDB *MongoInstance

func ConnectMongo() *MongoInstance {
	opts := options.Client().ApplyURI(config.MONGODB_ADDRESS["HOST"])
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database(config.MONGODB_ADDRESS["DATABASE"])
	instance := MongoInstance{
		Client: client,
		Db:     db,
	}
	return &instance
}

func init() {
	MongoDB = ConnectMongo()
	fmt.Println("CONNECT MONGODB SUCCESS!")
}
