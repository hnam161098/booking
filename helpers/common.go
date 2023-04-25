package helpers

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

func ConvertTimeFromES(t string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	tmp, err := time.Parse(layout, t)
	if err != nil {
		fmt.Println(err)
	}
	return tmp
}

func ValidateStruct(user interface{}) []*ErrorResponseValidate {
	var errors []*ErrorResponseValidate
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseValidate
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}
