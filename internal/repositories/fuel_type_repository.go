package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type FuelTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.FuelType, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.FuelType, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.FuelType, error)

	CreateRepository(
		ctx context.Context,
		FuelType *db_models.FuelType,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateFuelTypeRequest,
	) (*db_models.FuelType, error)
}
