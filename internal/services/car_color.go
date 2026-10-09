package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func IsEmptyColor(color *api_models.CarColor) bool {
	return color != nil &&
		color.Id == nil &&
		color.Code == nil
}

func ResolveColorID(
	ctx context.Context,
	color *api_models.CarColor,
) (*uuid.UUID, error) {
	if color == nil {
		return nil, nil
	}
	colorRepo := repositories.NewColorRepository()
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
			result := validation.New()
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
		newColor := &db_models.Color{
			Code: *color.Code,
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
