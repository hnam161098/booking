package helpers

var Errors map[string]int
var ErrorsCustomer map[string]int
var ErrorsBooking map[string]int

func init() {
	Errors = map[string]int{
		"MISSING_PARAMS": 100,
		"NOT_FOUND":      101,
	}

	ErrorsCustomer = map[string]int{
		"CREATE_CUSTOMER_ERROR":      1000,
		"UPDATE_CUSTOMER_ERROR":      1001,
		"FIND_CUSTOMER_ERROR":        1002,
		"DELETE_CUSTOMER_ERROR":      1003,
		"ADD_TAGS_CUSTOMER_ERROR":    1004,
		"DELETE_TAGS_CUSTOMER_ERROR": 1005,
	}

	ErrorsBooking = map[string]int{
		"CREATE_TICKET_ERROR": 2000,
		"FIND_TICKET_ERROR":   2001,
	}
}
