package helpers

type ResponseJSON struct {
	Status  bool
	Message string
	Error   string
	Payload interface{}
}

type ErrorResponse struct {
	ErrorCode    int
	ErrorMessage string
	ErrorSystem  string
}

type ErrorResponseValidate struct {
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value,omitempty"`
}
