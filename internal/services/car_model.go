package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)
func IsEmptyModel(
	model *api_models.CarModel,
) bool {
	return model != nil &&
		model.Id == nil &&
		model.Name == nil &&
		model.Brand == nil
}

func ResolveModelID(
	ctx context.Context,
	model *api_models.CarModel,
) (*uuid.UUID, error) {
	if model == nil {
		return nil, nil
	}
	modelRepo := repositories.NewModelRepository()
	if model.Id != nil {
		modelID, err := uuid.Parse(
			*model.Id,
		)
		if err == nil {
			existingModel, err := modelRepo.GetByIDRepository(
				ctx,
				modelID,
			)
			if err == nil {
				return &existingModel.ID, nil
			}
		}
		if model.Name == nil {
			result := validation.New()
			result.Add(
				"model",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if model.Name != nil {
		var brandID *uuid.UUID
		if model.Brand != nil &&
			!IsEmptyNameReference(model.Brand) {
			resolvedBrandID, err := ResolveBrandID(
				ctx,
				model.Brand,
			)
			if err != nil {
				return nil, err
			}
			brandID = resolvedBrandID
		}
		existingModel, err := modelRepo.GetByNameRepository(
			ctx,
			*model.Name,
			brandID,
		)
		if err == nil {
			return &existingModel.ID, nil
		}
		newModel := &db_models.Model{
			Name:    *model.Name,
			BrandID: brandID,
		}
		err = modelRepo.CreateRepository(
			ctx,
			newModel,
		)
		if err != nil {
			return nil, err
		}
		return &newModel.ID, nil
	}

	return nil, nil
}