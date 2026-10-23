package car_transmission

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func ResolveTransmissionTypeID(
	ctx context.Context,
	item *api_models.CodeReference,
) (*uuid.UUID, error) {
	if item == nil {
		return nil, nil
	}
	repo := NewTransmissionTypeRepository()
	if item.Id != nil {
		itemID, err := uuid.Parse(*item.Id)
		if err == nil {
			existItem, err := repo.GetTransmissionTypeByIDRepository(
				ctx,
				itemID,
			)
			if err == nil {
				return &existItem.ID, nil
			}
		}
		if item.Code == nil {
			result := shared_validation.NewValidationResult()
			result.Add(
				"bodyType",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if item.Code != nil {
		existItem, err := repo.GetTransmissionTypeByCodeRepository(
			ctx,
			*item.Code,
		)
		if err == nil {
			return &existItem.ID, nil
		}
		newItem := &TransmissionTypeDBModel{
			Code: *item.Code,
		}
		err = repo.CreateTransmissionTypeRepository(
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
