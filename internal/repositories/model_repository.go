package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type ModelRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.Model, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.Model, error)

	GetByNameRepository(
		ctx context.Context,
		name string,
	) (*db_models.Model, error)

	CreateRepository(
		ctx context.Context,
		Model *db_models.Model,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateModelRequest,
	) (*db_models.Model, error)
}
