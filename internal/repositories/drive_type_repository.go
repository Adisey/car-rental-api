package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type DriveTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.DriveType, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.DriveType, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.DriveType, error)

	CreateRepository(
		ctx context.Context,
		BodyType *db_models.DriveType,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateDriveTypeRequest,
	) (*db_models.DriveType, error)
}
