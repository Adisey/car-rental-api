package car_brand

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type BrandRepository interface {
	GetBrandAllRepository(ctx context.Context) ([]BrandDBModel, error)

	GetBrandByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*BrandDBModel, error)

	GetBrandByNameRepository(
		ctx context.Context,
		name string,
	) (*BrandDBModel, error)

	CreateBrandRepository(
		ctx context.Context,
		Brand *BrandDBModel,
	) error

	UpdateBrandRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBrandRequest,
	) (*BrandDBModel, error)
}
