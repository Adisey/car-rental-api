package car_model

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type ModelRepository interface {
	GetModelAllRepository(ctx context.Context) ([]ModelDBModel, error)

	GetModelByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*ModelDBModel, error)

	GetModelByNameRepository(
		ctx context.Context,
		name string,
	) (*ModelDBModel, error)

	CreateModelRepository(
		ctx context.Context,
		Model *ModelDBModel,
	) error

	UpdateModelRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateModelRequest,
	) (*ModelDBModel, error)
}
