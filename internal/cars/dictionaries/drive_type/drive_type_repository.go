package car_drive_types

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type DriveTypeRepository interface {
	GetDriveTypeAllRepository(ctx context.Context) ([]DriveTypeDBModel, error)

	GetDriveTypeByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*DriveTypeDBModel, error)

	GetDriveTypeByCodeRepository(
		ctx context.Context,
		code string,
	) (*DriveTypeDBModel, error)

	CreateDriveTypeRepository(
		ctx context.Context,
		BodyType *DriveTypeDBModel,
	) error

	UpdateDriveTypeRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateDriveTypeRequest,
	) (*DriveTypeDBModel, error)
}
