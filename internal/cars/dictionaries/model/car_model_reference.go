package car_model

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	shared_references "github.com/Adisey/car-rental-api/internal/shared/references"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func ResolveModelID(
	ctx context.Context,
	model *api_models.CarModel,
) (*uuid.UUID, error) {
	if model == nil {
		return nil, nil
	}
	modelRepo := NewModelRepository()
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
			result := shared_validation.NewValidationResult()
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
			!shared_references.IsEmptyNameReference(model.Brand) {
			resolvedBrandID, err := car_brand.ResolveBrandID(
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
		newModel := &ModelDBModel{
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
