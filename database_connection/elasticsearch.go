package database_connection

import (
	"grpc/config"
	"log"

	"github.com/elastic/go-elasticsearch"
)

var ESconnection *elasticsearch.Client

func ConnectES() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{config.ELASTICSEARCH_ADDRESS["ENDPOINT"]},
	}
	ESConn, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("CONNECTING ES ERROR")
	}
	if ESConn == nil {
		log.Println("CONNECTING ES ERROR")
	}
	return ESConn
}

func init() {
	ESconnection = ConnectES()
}
