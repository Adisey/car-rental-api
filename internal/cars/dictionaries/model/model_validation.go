package car_model

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

var ModelRules = struct {
	Name shared_validation.StringRule
}{
	Name: shared_validation.StringRule{
		MinLength: 2,
		MaxLength: 100,
	},
}

func ValidateCreateModelRequest(
	request api_models.CreateModelRequest,
) *shared_validation.ValidationResult {

	result := shared_validation.NewValidationResult()

	name := strings.TrimSpace(request.Name)

	if name == "" {
		result.Add(
			"name",
			"required",
			nil,
		)
	}

	if len(name) < ModelRules.Name.MinLength {
		result.Add(
			"name",
			"min_length",
			map[string]any{
				"min": ModelRules.Name.MinLength,
			},
		)
	}

	if len(name) > ModelRules.Name.MaxLength {
		result.Add(
			"name",
			"max_length",
			map[string]any{
				"max": ModelRules.Name.MaxLength,
			},
		)
	}

	return result
}

func ValidateUpdateModelRequest(
	request api_models.UpdateModelRequest,
) *shared_validation.ValidationResult {

	result := shared_validation.NewValidationResult()

	if request.Name != nil {
		name := strings.TrimSpace(*request.Name)

		if name == "" {
			result.Add(
				"name",
				"required",
				nil,
			)
		}

		if len(name) < ModelRules.Name.MinLength {
			result.Add(
				"name",
				"min_length",
				map[string]any{
					"min": ModelRules.Name.MinLength,
				},
			)
		}

		if len(name) > ModelRules.Name.MaxLength {
			result.Add(
				"name",
				"max_length",
				map[string]any{
					"max": ModelRules.Name.MaxLength,
				},
			)
		}

	}

	return result
}
