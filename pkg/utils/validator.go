package utils

import (
	"github.com/go-playground/validator/v10"
	"summershare/internal/entities/web"
)

var validate = validator.New()

func ValidateStruct(body interface{}) []*web.ValidationErrorResponse {
	var errors []*web.ValidationErrorResponse
	if err := validate.Struct(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &web.ValidationErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
				Message:     err.Error(),
			})
		}
	}

	return errors
}
