package car

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type CarRepository interface {
	GetAllRepository(ctx context.Context) ([]CarDBModel, error)
	GetByIDRepository(ctx context.Context, id uuid.UUID) (*CarDBModel, error)
	CreateRepository(ctx context.Context, car *CarDBModel) error
	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateCarRequest,
		colorID *uuid.UUID,
		shouldUpdateColor bool,
	) (*CarDBModel, error)
	DeleteRepository(
		ctx context.Context,
		id uuid.UUID,
	) error
}
