package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BodyTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.BodyType, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.BodyType, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.BodyType, error)

	CreateRepository(
		ctx context.Context,
		BodyType *db_models.BodyType,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBodyTypeRequest,
	) (*db_models.BodyType, error)
}
