package validation

import (
	"github.com/Adisey/car-rental-api/internal/api_models"
)

type ValidationResult struct {
	fields map[string][]api_models.ValidationError
}

func New() *ValidationResult {
	return &ValidationResult{
		fields: make(map[string][]api_models.ValidationError),
	}
}

func (v *ValidationResult) Add(
	field string,
	code string,
	params map[string]any,
) {
	v.fields[field] = append(
		v.fields[field],
		api_models.ValidationError{
			Code:   code,
			Params: params,
		},
	)
}

func (v *ValidationResult) HasErrors() bool {
	return len(v.fields) > 0
}

func (v *ValidationResult) Response() api_models.ValidationErrorsResponse {
	return api_models.ValidationErrorsResponse{
		Errors: v.fields,
	}
}

func (v *ValidationResult) Error() string {
	return "validation failed"
}


func AddError(
	errors map[string][]api_models.ValidationError,
	field string,
	code string,
	params map[string]any,
) {
	errors[field] = append(
		errors[field],
		api_models.ValidationError{
			Code:   code,
			Params: params,
		},
	)
}
