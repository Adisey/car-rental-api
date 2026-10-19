package car_transmission_type

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type TransmissionTypeRepository interface {
	GetAllRepository(ctx context.Context) ([]TransmissionTypeDBModel, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*TransmissionTypeDBModel, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*TransmissionTypeDBModel, error)

	CreateRepository(
		ctx context.Context,
		TransmissionType *TransmissionTypeDBModel,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateTransmissionTypeRequest,
	) (*TransmissionTypeDBModel, error)
}
