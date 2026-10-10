package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BrandRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.Brand, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.Brand, error)

	GetByNameRepository(
		ctx context.Context,
		name string,
	) (*db_models.Brand, error)

	CreateRepository(
		ctx context.Context,
		Brand *db_models.Brand,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBrandRequest,
	) (*db_models.Brand, error)
}
