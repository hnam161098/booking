package config

var ADDRESS_SERVER map[string]string
var ADDRESS_CLIENT map[string]string
var MONGODB_ADDRESS map[string]string
var ELASTICSEARCH_ADDRESS map[string]string

func init() {
	ADDRESS_SERVER = map[string]string{
		"CUSTOMER_PORT": ":10001",
		"BOOKING_PORT":  ":10002",
	}
	ADDRESS_CLIENT = map[string]string{
		"CUSTOMER_PORT": ":20001",
		"BOOKING_PORT":  ":20002",
	}
	MONGODB_ADDRESS = map[string]string{
		"HOST":     "mongodb://localhost",
		"PORT":     "27017",
		"DATABASE": "data_dev",
	}

	ELASTICSEARCH_ADDRESS = map[string]string{
		"ENDPOINT": "http://localhost:9200",
	}
}
