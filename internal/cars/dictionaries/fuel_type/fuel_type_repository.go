package car_fuel

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type FuelTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]FuelTypeDBModel, error)

	GetFuelTypeByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*FuelTypeDBModel, error)

	GetFuelTypeByCodeRepository(
		ctx context.Context,
		code string,
	) (*FuelTypeDBModel, error)

	CreateFuelTypeRepository(
		ctx context.Context,
		FuelType *FuelTypeDBModel,
	) error

	UpdateFuelTypeRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateFuelTypeRequest,
	) (*FuelTypeDBModel, error)
}
