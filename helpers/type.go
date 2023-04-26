package helpers

type ResponseJSON struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ErrorSystem  string `json:"error_system"`
}

type ErrorResponseValidate struct {
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value,omitempty"`
}
