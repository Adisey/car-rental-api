package validation

import (
	"strings"

	api_models "github.com/Adisey/car-rental-api/internal/api_models"
)

var ColorRules = struct {
	Code        StringRule
	Description StringRule
}{
	Code: StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
}

func ValidateCreateColorRequest(
	request api_models.CreateColorRequest,
) *ValidationResult {

	result := New()

	code := strings.TrimSpace(request.Code)

	if code == "" {
		result.Add(
			"name",
			"required",
			nil,
		)
	}

	if len(code) < ColorRules.Code.MinLength {
		result.Add(
			"name",
			"min_length",
			map[string]any{
				"min": ColorRules.Code.MinLength,
			},
		)
	}

	if len(code) > ColorRules.Code.MaxLength {
		result.Add(
			"name",
			"max_length",
			map[string]any{
				"max": ColorRules.Code.MaxLength,
			},
		)
	}

	return result
}

func ValidateUpdateColorRequest(
	request api_models.UpdateColorRequest,
) *ValidationResult {

	result := New()

	if request.Code != nil {
		code := strings.TrimSpace(*request.Code)

		if code == "" {
			result.Add(
				"code",
				"required",
				nil,
			)
		}

		if len(code) < ColorRules.Code.MinLength {
			result.Add(
				"name",
				"min_length",
				map[string]any{
					"min": ColorRules.Code.MinLength,
				},
			)
		}
	}

	return result
}
