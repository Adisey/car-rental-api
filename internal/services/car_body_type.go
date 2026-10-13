package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func IsEmptyBodyType(bodyType *api_models.CodeReference) bool {
	return bodyType != nil &&
		bodyType.Id == nil &&
		bodyType.Code == nil
}

func ResolveBodyTypeID(
	ctx context.Context,
	bodyType *api_models.CodeReference,
) (*uuid.UUID, error) {
	if bodyType == nil {
		return nil, nil
	}
	bodyTypeRepo := repositories.NewBodyTypeRepository()
	if bodyType.Id != nil {
		bodyTypeID, err := uuid.Parse(*bodyType.Id)
		if err == nil {
			existingBodyType, err := bodyTypeRepo.GetByIDRepository(
				ctx,
				bodyTypeID,
			)
			if err == nil {
				return &existingBodyType.ID, nil
			}
		}
		if bodyType.Code == nil {
			result := validation.New()
			result.Add(
				"bodyType",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if bodyType.Code != nil {
		existingBodyType, err := bodyTypeRepo.GetByCodeRepository(
			ctx,
			*bodyType.Code,
		)
		if err == nil {
			return &existingBodyType.ID, nil
		}
		newBodyType := &db_models.BodyType{
			Code: *bodyType.Code,
		}
		err = bodyTypeRepo.CreateRepository(
			ctx,
			newBodyType,
		)
		if err != nil {
			return nil, err
		}
		return &newBodyType.ID, nil
	}
	return nil, nil
}
