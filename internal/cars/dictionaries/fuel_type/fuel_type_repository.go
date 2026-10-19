package car_fuel

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type FuelTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]FuelTypeDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*FuelTypeDBModel, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*FuelTypeDBModel, error)

	CreateRepository(
		ctx context.Context,
		FuelType *FuelTypeDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateFuelTypeRequest,
	) (*FuelTypeDBModel, error)
}
