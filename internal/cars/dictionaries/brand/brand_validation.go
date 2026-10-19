package car_brand

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

var BrandRules = struct {
	Name shared_validation.StringRule
}{
	Name: shared_validation.StringRule{
		MinLength: 2,
		MaxLength: 100,
	},
}

func ValidateCreateBrandRequest(
	request api_models.CreateBrandRequest,
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

	}

	return result
}
