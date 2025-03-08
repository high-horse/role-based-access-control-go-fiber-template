package utils

import (
	// "fmt"
	// "strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Error       bool        `json:"error"`
	FailedField string      `json:"failed_field"`
	Tag         string      `json:"tag"`
	Value       interface{} `json:"value"`
}

type XValidator struct {
	validator *validator.Validate
}

// Initialize the validator instance
var validate = validator.New()

// Register custom validations (if needed)
func RegisterCustomValidations() {
	// Example of a custom validator for "teener" age validation (12-18 years)
	validate.RegisterValidation("teener", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
	})
}

// Validate function to validate any struct
func (v *XValidator) Validate(data interface{}) ([]ErrorResponse, error) {
	var validationErrors []ErrorResponse

	// Perform validation using the go-playground validator
	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// Creating a structured error response
			validationErrors = append(validationErrors, ErrorResponse{
				Error:       true,
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}

	return validationErrors, nil
}

// ValidateRequest function that can be used directly for any request validation
func ValidateRequest(data interface{}) ([]ErrorResponse, error) {
	// Create a new instance of XValidator
	validator := &XValidator{validator: validate}

	// Validate the data using the helper function
	return validator.Validate(data)
}
