package car_body_type

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type BodyTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]BodyTypeDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*BodyTypeDBModel, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*BodyTypeDBModel, error)

	CreateRepository(
		ctx context.Context,
		BodyType *BodyTypeDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBodyTypeRequest,
	) (*BodyTypeDBModel, error)
}
