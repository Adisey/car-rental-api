package car_model

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type ModelRepository interface {
	GetAllRepository(ctx context.Context) ([]ModelDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*ModelDBModel, error)

	GetByNameRepository(
		ctx context.Context,
		name string,
	) (*ModelDBModel, error)

	CreateRepository(
		ctx context.Context,
		Model *ModelDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateModelRequest,
	) (*ModelDBModel, error)
}
