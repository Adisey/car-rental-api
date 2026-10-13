package validation

import (
	"strings"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

var BodyTypeRules = struct {
	Code        StringRule
	Description StringRule
}{
	Code: StringRule{
		MinLength: 3,
		MaxLength: 100,
	},
}

func ValidateCreateBodyTypeRequest(
	request api_models.CreateBodyTypeRequest,
) *ValidationResult {

	result := New()

	code := strings.TrimSpace(request.Code)

	if code == "" {
		result.Add(
			"code",
			"required",
			nil,
		)
	}

	if len(code) < BodyTypeRules.Code.MinLength {
		result.Add(
			"code",
			"min_length",
			map[string]any{
				"min": BodyTypeRules.Code.MinLength,
			},
		)
	}

	if len(code) > BodyTypeRules.Code.MaxLength {
		result.Add(
			"code",
			"max_length",
			map[string]any{
				"max": BodyTypeRules.Code.MaxLength,
			},
		)
	}

	return result
}

func ValidateUpdateBodyTypeRequest(
	request api_models.UpdateBodyTypeRequest,
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

		if len(code) < BodyTypeRules.Code.MinLength {
			result.Add(
				"code",
				"min_length",
				map[string]any{
					"min": BodyTypeRules.Code.MinLength,
				},
			)
		}

		if len(code) > BodyTypeRules.Code.MaxLength {
			result.Add(
				"code",
				"max_length",
				map[string]any{
					"max": BodyTypeRules.Code.MaxLength,
				},
			)
		}

	}

	return result
}
