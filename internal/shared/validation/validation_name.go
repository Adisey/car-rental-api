package shared_validation

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

var NameEntityRules = struct{ Name StringRule }{
	Name: StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
}

func ValidateCreateNameRequest(request api_models.CreateNameEntityRequest) *ValidationResult {
	result := NewValidationResult()
	name := strings.TrimSpace(request.Name)
	if name == "" {
		result.Add("name", "required", nil)
	}
	if len(name) < NameEntityRules.Name.MinLength {
		result.Add(
			"name",
			"min_length",
			map[string]any{
				"min": NameEntityRules.Name.MinLength,
			},
		)
	}
	if len(name) > NameEntityRules.Name.MaxLength {
		result.Add(
			"name",
			"max_length",
			map[string]any{
				"max": NameEntityRules.Name.MaxLength,
			},
		)
	}
	return result
}

func ValidateUpdateNameRequest(request api_models.UpdateNameEntityRequest) *ValidationResult {
	result := NewValidationResult()
	if request.Name != nil {
		name := strings.TrimSpace(*request.Name)
		if name == "" {
			result.Add(
				"name",
				"required",
				nil,
			)
		}
		if len(name) < NameEntityRules.Name.MinLength {
			result.Add(
				"name",
				"min_length",
				map[string]any{
					"min": NameEntityRules.Name.MinLength,
				},
			)
		}
		if len(name) > NameEntityRules.Name.MaxLength {
			result.Add(
				"name",
				"max_length",
				map[string]any{
					"max": NameEntityRules.Name.MaxLength,
				},
			)
		}
	}
	return result
}
