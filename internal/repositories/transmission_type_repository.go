package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type TransmissionTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.TransmissionType, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.TransmissionType, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.TransmissionType, error)

	CreateRepository(
		ctx context.Context,
		TransmissionType *db_models.TransmissionType,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateTransmissionTypeRequest,
	) (*db_models.TransmissionType, error)
}
