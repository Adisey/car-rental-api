package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func ResolveFuelTypeID(
	ctx context.Context,
	item *api_models.CodeReference,
) (*uuid.UUID, error) {
	if item == nil {
		return nil, nil
	}
	repo := repositories.NewFuelTypeRepository()
	if item.Id != nil {
		itemID, err := uuid.Parse(*item.Id)
		if err == nil {
			existItem, err := repo.GetByIDRepository(
				ctx,
				itemID,
			)
			if err == nil {
				return &existItem.ID, nil
			}
		}
		if item.Code == nil {
			result := validation.New()
			result.Add(
				"bodyType",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if item.Code != nil {
		existItem, err := repo.GetByCodeRepository(
			ctx,
			*item.Code,
		)
		if err == nil {
			return &existItem.ID, nil
		}
		newItem := &db_models.FuelType{
			Code: *item.Code,
		}
		err = repo.CreateRepository(
			ctx,
			newItem,
		)
		if err != nil {
			return nil, err
		}
		return &newItem.ID, nil
	}
	return nil, nil
}
