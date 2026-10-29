package customers

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

var CustomerRules = struct {
	FirstName      shared_validation.StringRule
	OptionsText100 shared_validation.StringRule
	OptionsText50  shared_validation.StringRule
}{
	FirstName: shared_validation.StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
	OptionsText100: shared_validation.StringRule{
		MaxLength: 100,
	},
	OptionsText50: shared_validation.StringRule{
		MaxLength: 50,
	},
}

func ValidateCreateCustomerRequest(
	request api_models.CreateCustomerRequest,
) *shared_validation.ValidationResult {
	result := shared_validation.NewValidationResult()
	firstName := strings.TrimSpace(request.FirstName)
	if firstName == "" {
		result.Add("first_name", "required", nil)
	}
	if len(firstName) < CustomerRules.FirstName.MinLength {
		result.Add("first_name", "min_length", map[string]any{"min": CustomerRules.FirstName.MinLength})
	}
	if len(firstName) > CustomerRules.FirstName.MaxLength {
		result.Add("first_name", "max_length", map[string]any{"max": CustomerRules.FirstName.MaxLength})
	}
	if request.LastName != nil {
		lastName := strings.TrimSpace(*request.LastName)
		if len(lastName) > CustomerRules.OptionsText100.MaxLength {
			result.Add("last_name", "max_length", map[string]any{"max": CustomerRules.OptionsText100.MaxLength})
		}
	}
	if request.Phone != nil {
		phone := strings.TrimSpace(*request.Phone)
		if len(phone) > CustomerRules.OptionsText50.MaxLength {
			result.Add("phone", "max_length", map[string]any{"max": CustomerRules.OptionsText50.MaxLength})
		}
	}
	if request.Email != nil {
		email := strings.TrimSpace(*request.Email)
		if len(email) > CustomerRules.OptionsText100.MaxLength {
			result.Add("email", "max_length", map[string]any{"max": CustomerRules.OptionsText100.MaxLength})
		}
	}
	return result
}

func ValidateUpdateCustomerRequest(
	request api_models.UpdateCustomerRequest,
) *shared_validation.ValidationResult {
	result := shared_validation.NewValidationResult()
	if request.FirstName != nil {
		firstName := strings.TrimSpace(*request.FirstName)
		if len(firstName) < CustomerRules.FirstName.MinLength {
			result.Add("first_name", "min_length", map[string]any{"min": CustomerRules.FirstName.MinLength})
		}
		if len(firstName) > CustomerRules.FirstName.MaxLength {
			result.Add("first_name", "max_length", map[string]any{"max": CustomerRules.FirstName.MaxLength})
		}
	}
	if request.LastName != nil {
		lastName := strings.TrimSpace(*request.LastName)
		if len(lastName) > CustomerRules.OptionsText100.MaxLength {
			result.Add("last_name", "max_length", map[string]any{"max": CustomerRules.OptionsText100.MaxLength})
		}
	}
	if request.Phone != nil {
		phone := strings.TrimSpace(*request.Phone)
		if len(phone) > CustomerRules.OptionsText50.MaxLength {
			result.Add("phone", "max_length", map[string]any{"max": CustomerRules.OptionsText50.MaxLength})
		}
	}
	if request.Email != nil {
		email := strings.TrimSpace(*request.Email)
		if len(email) > CustomerRules.OptionsText100.MaxLength {
			result.Add("email", "max_length", map[string]any{"max": CustomerRules.OptionsText100.MaxLength})
		}
	}
	return result
}
