package car_color

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func ResolveColorID(
	ctx context.Context,
	color *api_models.CarColor,
) (*uuid.UUID, error) {
	if color == nil {
		return nil, nil
	}
	colorRepo := NewColorRepository()
	if color.Id != nil {
		colorID, err := uuid.Parse(*color.Id)
		if err == nil {
			existingColor, err := colorRepo.GetByIDRepository(
				ctx,
				colorID,
			)
			if err == nil {
				return &existingColor.ID, nil
			}
		}
		if color.Code == nil {
			result := shared_validation.NewValidationResult()
			result.Add(
				"color",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if color.Code != nil {
		existingColor, err := colorRepo.GetByCodeRepository(
			ctx,
			*color.Code,
		)
		if err == nil {
			return &existingColor.ID, nil
		}
		newColor := &ColorDBModel{
			Code: *color.Code,
			Hex:  color.Hex,
		}
		err = colorRepo.CreateRepository(
			ctx,
			newColor,
		)
		if err != nil {
			return nil, err
		}
		return &newColor.ID, nil
	}
	return nil, nil
}
