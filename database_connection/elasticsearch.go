package database_connection

import (
	"fmt"
	"grpc/config"

	"github.com/elastic/go-elasticsearch"
)

var ESconnection *elasticsearch.Client

func ConnectES() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{config.ELASTICSEARCH_ADDRESS["ENDPOINT"]},
	}
	ESConn, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("CONNECTING ES ERROR")
	}
	return ESConn
}

func init() {
	ESconnection = ConnectES()
	fmt.Println("CONNECT ELASTICSEARCH SUCCESS!")
}
