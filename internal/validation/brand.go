package validation

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

var BrandRules = struct {
	Name        StringRule
	Description StringRule
}{
	Name: StringRule{
		MinLength: 2,
		MaxLength: 100,
	},
}

func ValidateCreateBrandRequest(
	request api_models.CreateBrandRequest,
) *ValidationResult {

	result := New()

	name := strings.TrimSpace(request.Name)

	if name == "" {
		result.Add(
			"name",
			"required",
			nil,
		)
	}

	if len(name) < BrandRules.Name.MinLength {
		result.Add(
			"name",
			"min_length",
			map[string]any{
				"min": BrandRules.Name.MinLength,
			},
		)
	}

	if len(name) > BrandRules.Name.MaxLength {
		result.Add(
			"name",
			"max_length",
			map[string]any{
				"max": BrandRules.Name.MaxLength,
			},
		)
	}

	return result
}

func ValidateUpdateBrandRequest(
	request api_models.UpdateBrandRequest,
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

		if len(name) < BrandRules.Name.MinLength {
			result.Add(
				"name",
				"min_length",
				map[string]any{
					"min": BrandRules.Name.MinLength,
				},
			)
		}
	}

	return result
}
