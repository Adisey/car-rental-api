package car_color

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type ColorRepository interface {
	GetAllRepository(ctx context.Context) ([]ColorDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*ColorDBModel, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*ColorDBModel, error)

	CreateRepository(
		ctx context.Context,
		color *ColorDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateColorRequest,
	) (*ColorDBModel, error)
}
