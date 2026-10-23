package car_brand

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func ResolveBrandID(
	ctx context.Context,
	brand *api_models.NameReference,
) (*uuid.UUID, error) {
	if brand == nil {
		return nil, nil
	}
	brandRepo := NewBrandRepository()
	if brand.Id != nil {
		brandID, err := uuid.Parse(*brand.Id)
		if err == nil {
			existingBrand, err := brandRepo.GetBrandByIDRepository(
				ctx,
				brandID,
			)
			if err == nil {
				return &existingBrand.ID, nil
			}
		}
		if brand.Name == nil {
			result := shared_validation.NewValidationResult()
			result.Add(
				"brand",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if brand.Name != nil {
		existingBrand, err := brandRepo.GetBrandByNameRepository(
			ctx,
			*brand.Name,
		)
		if err == nil {
			return &existingBrand.ID, nil
		}
		newBrand := &BrandDBModel{
			Name: *brand.Name,
		}
		err = brandRepo.CreateBrandRepository(
			ctx,
			newBrand,
		)
		if err != nil {
			return nil, err
		}
		return &newBrand.ID, nil
	}
	return nil, nil
}
