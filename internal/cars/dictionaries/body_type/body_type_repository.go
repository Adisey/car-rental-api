package car_body

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type BodyTypeRepository interface {
	GetBodyTypeAllRepository(ctx context.Context) ([]BodyTypeDBModel, error)

	GetBodyTypeByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*BodyTypeDBModel, error)

	GetBodyTypeByCodeRepository(
		ctx context.Context,
		code string,
	) (*BodyTypeDBModel, error)

	CreateBodyTypeRepository(
		ctx context.Context,
		BodyType *BodyTypeDBModel,
	) error

	UpdateBodyTypeRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateBodyTypeRequest,
	) (*BodyTypeDBModel, error)
}
