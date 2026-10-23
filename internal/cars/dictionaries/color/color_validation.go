package car_color

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

// var ColorRules = struct {
// 	Code shared_validation.StringRule
// 	Hex  shared_validation.StringRule
// }{
// 	Code: shared_validation.StringRule{
// 		MinLength: 3,
// 		MaxLength: 100,
// 	},
// }

func ValidateCreateColorRequest(
	request api_models.CreateColorRequest,
) *shared_validation.ValidationResult {

	result := shared_validation.NewValidationResult()

	code := strings.TrimSpace(request.Code)

	if code == "" {
		result.Add(
			"code",
			"required",
			nil,
		)
	}

	if len(code) < shared_validation.CodeEntityRules.Code.MinLength {
		result.Add(
			"code",
			"min_length",
			map[string]any{
				"min": shared_validation.CodeEntityRules.Code.MinLength,
			},
		)
	}

	if len(code) > shared_validation.CodeEntityRules.Code.MaxLength {
		result.Add(
			"code",
			"max_length",
			map[string]any{
				"max": shared_validation.CodeEntityRules.Code.MaxLength,
			},
		)
	}

	return result
}

func ValidateUpdateColorRequest(
	request api_models.UpdateColorRequest,
) *shared_validation.ValidationResult {

	result := shared_validation.NewValidationResult()

	if request.Code != nil {
		code := strings.TrimSpace(*request.Code)

		if code == "" {
			result.Add(
				"code",
				"required",
				nil,
			)
		}

		if len(code) < shared_validation.CodeEntityRules.Code.MinLength {
			result.Add(
				"code",
				"min_length",
				map[string]any{
					"min": shared_validation.CodeEntityRules.Code.MinLength,
				},
			)
		}

		if len(code) > shared_validation.CodeEntityRules.Code.MaxLength {
			result.Add(
				"code",
				"max_length",
				map[string]any{
					"max": shared_validation.CodeEntityRules.Code.MaxLength,
				},
			)
		}

	}

	return result
}
