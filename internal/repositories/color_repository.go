package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type ColorRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.Color, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.Color, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.Color, error)

	CreateRepository(
		ctx context.Context,
		color *db_models.Color,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateColorRequest,
	) (*db_models.Color, error)
}
