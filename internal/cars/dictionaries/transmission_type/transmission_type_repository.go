package car_transmission

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type TransmissionTypeRepository interface {
	GetTransmissionTypeAllRepository(ctx context.Context) ([]TransmissionTypeDBModel, error)

	GetTransmissionTypeByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*TransmissionTypeDBModel, error)

	GetTransmissionTypeByCodeRepository(
		ctx context.Context,
		code string,
	) (*TransmissionTypeDBModel, error)

	CreateTransmissionTypeRepository(
		ctx context.Context,
		TransmissionType *TransmissionTypeDBModel,
	) error

	UpdateTransmissionTypeRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateTransmissionTypeRequest,
	) (*TransmissionTypeDBModel, error)
}
