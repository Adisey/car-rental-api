package car_brand

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type BrandRepository interface {
	GetAllRepository(ctx context.Context) ([]BrandDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*BrandDBModel, error)

	GetByNameRepository(
		ctx context.Context,
		name string,
	) (*BrandDBModel, error)

	CreateRepository(
		ctx context.Context,
		Brand *BrandDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBrandRequest,
	) (*BrandDBModel, error)
}
