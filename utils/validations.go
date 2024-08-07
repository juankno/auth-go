package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

var validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validate.Struct(data)
}

func FormatValidationError(err error) ErrorResponse {
	var errors []ValidationError

	// Verifica si el error es de tipo ValidationErrors
	if _, ok := err.(validator.ValidationErrors); ok {
		for _, err := range err.(validator.ValidationErrors) {
			// Crea un error personalizado
			ve := ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errors = append(errors, ve)
		}
	}

	return ErrorResponse{Errors: errors}
}
