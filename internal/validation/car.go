package validation

import (
	"strings"

	api_models "github.com/Adisey/car-rental-api/internal/api_models"
)

var CarRules = struct {
	Name        StringRule
	Description StringRule
}{
	Name: StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
	Description: StringRule{
		MaxLength: 1000,
	},
}

func ValidateCreateCarRequest(
	request api_models.CreateCarRequest,
) *ValidationResult {

	result := New()

	name := strings.TrimSpace(request.Name)
	description := strings.TrimSpace(*request.Description)

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

	if len(description) > CarRules.Description.MaxLength {
		result.Add(
			"description",
			"max_length",
			map[string]any{
				"max": CarRules.Description.MaxLength,
			},
		)
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

func ValidateUpdateCarRequest(
	request api_models.UpdateCarRequest,
) *ValidationResult {

	result := New()

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
