package validation

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

var CodeEntityRules = struct{ Code StringRule }{
	Code: StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
}

func ValidateCreateCodeRequest(request api_models.CreateCodeEntityRequest) *ValidationResult {
	result := New()
	code := strings.TrimSpace(request.Code)
	if code == "" {
		result.Add("code", "required", nil)
	}
	if len(code) < CodeEntityRules.Code.MinLength {
		result.Add(
			"code",
			"min_length",
			map[string]any{
				"min": CodeEntityRules.Code.MinLength,
			},
		)
	}
	if len(code) > CodeEntityRules.Code.MaxLength {
		result.Add(
			"code",
			"max_length",
			map[string]any{
				"max": CodeEntityRules.Code.MaxLength,
			},
		)
	}
	return result
}

func ValidateUpdateCodeRequest(request api_models.UpdateCodeEntityRequest) *ValidationResult {
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
		if len(code) < CodeEntityRules.Code.MinLength {
			result.Add(
				"code",
				"min_length",
				map[string]any{
					"min": CodeEntityRules.Code.MinLength,
				},
			)
		}
		if len(code) > CodeEntityRules.Code.MaxLength {
			result.Add(
				"code",
				"max_length",
				map[string]any{
					"max": CodeEntityRules.Code.MaxLength,
				},
			)
		}
	}
	return result
}
