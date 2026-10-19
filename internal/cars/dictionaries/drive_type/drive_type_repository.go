package car_drive_types

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type DriveTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]DriveTypeDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*DriveTypeDBModel, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*DriveTypeDBModel, error)

	CreateRepository(
		ctx context.Context,
		BodyType *DriveTypeDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateDriveTypeRequest,
	) (*DriveTypeDBModel, error)
}
