package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func IsEmptyBrand(brand *api_models.CarBrand) bool {
	return brand != nil &&
		brand.Id == nil &&
		brand.Name == nil
}

func ResolveBrandID(
	ctx context.Context,
	brand *api_models.CarBrand,
) (*uuid.UUID, error) {
	if brand == nil {
		return nil, nil
	}
	brandRepo := repositories.NewBrandRepository()
	if brand.Id != nil {
		brandID, err := uuid.Parse(*brand.Id)
		if err == nil {
			existingBrand, err := brandRepo.GetByIDRepository(
				ctx,
				brandID,
			)
			if err == nil {
				return &existingBrand.ID, nil
			}
		}
		if brand.Name == nil {
			result := validation.New()
			result.Add(
				"brand",
				"not_found",
				nil,
			)
			return nil, result
		}
	}
	if brand.Name != nil {
		existingBrand, err := brandRepo.GetByNameRepository(
			ctx,
			*brand.Name,
		)
		if err == nil {
			return &existingBrand.ID, nil
		}
		newBrand := &db_models.Brand{
			Name: *brand.Name,
		}
		err = brandRepo.CreateRepository(
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
