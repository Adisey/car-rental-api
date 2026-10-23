package car_color

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type ColorRepository interface {
	GetColorAllRepository(ctx context.Context) ([]ColorDBModel, error)

	GetColorByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*ColorDBModel, error)

	GetColorByCodeRepository(
		ctx context.Context,
		code string,
	) (*ColorDBModel, error)

	CreateColorRepository(
		ctx context.Context,
		color *ColorDBModel,
	) error

	UpdateColorRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateColorRequest,
	) (*ColorDBModel, error)
}
