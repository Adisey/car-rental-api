package car

import (
	"strings"

	api_models "github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

var CarRules = struct {
	Name        shared_validation.StringRule
	Description shared_validation.StringRule
}{
	Name: shared_validation.StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
	Description: shared_validation.StringRule{
		MaxLength: 1000,
	},
}

func ValidateCreateCarRequest(
	request api_models.CreateCarRequest,
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
	if len(name) < CarRules.Name.MinLength {
		result.Add(
			"name",
			"min_length",
			map[string]any{
				"min": CarRules.Name.MinLength,
			},
		)
	}
	if len(name) > CarRules.Name.MaxLength {
		result.Add(
			"name",
			"max_length",
			map[string]any{
				"max": CarRules.Name.MaxLength,
			},
		)
	}
	description := ""
	if request.Description != nil {
		description = strings.TrimSpace(
			*request.Description,
		)
	}
	if len(description) > CarRules.Description.MaxLength {
		result.Add(
			"description",
			"max_length",
			map[string]any{
				"max": CarRules.Description.MaxLength,
			},
		)
	}
	if strings.Contains(
		strings.ToLower(description),
		"test",
	) {

		result.Add(
			"_object",
			"not_time_for_tests",
			nil,
		)
	}
	return result
}

func ValidateUpdateCarRequest(
	request api_models.UpdateCarRequest,
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

		if len(name) < CarRules.Name.MinLength {
			result.Add(
				"name",
				"min_length",
				map[string]any{
					"min": CarRules.Name.MinLength,
				},
			)
		}
	}

	if request.Description != nil &&
		strings.Contains(
			strings.ToLower(*request.Description),
			"test",
		) {

		result.Add(
			"_object",
			"not_time_for_tests",
			nil,
		)
	}

	return result
}
