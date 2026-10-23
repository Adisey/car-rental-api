package car

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type CarRepository interface {
	GetCarAllRepository(ctx context.Context) ([]CarDBModel, error)
	GetCarByIDRepository(ctx context.Context, id uuid.UUID) (*CarDBModel, error)
	CreateCarRepository(ctx context.Context, car *CarDBModel) error
	UpdateCarRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateCarRequest,
		colorID *uuid.UUID,
		shouldUpdateColor bool,
	) (*CarDBModel, error)
	DeleteCarRepository(
		ctx context.Context,
		id uuid.UUID,
	) error
}
