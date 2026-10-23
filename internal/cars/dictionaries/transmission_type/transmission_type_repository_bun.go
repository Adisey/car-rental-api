package car_transmission

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

type BunTransmissionTypeRepository struct{}

func NewTransmissionTypeRepository() *BunTransmissionTypeRepository {
	return &BunTransmissionTypeRepository{}
}

func (r *BunTransmissionTypeRepository) GetTransmissionTypeAllRepository(ctx context.Context) ([]TransmissionTypeDBModel, error) {
	var transmissionTypes []TransmissionTypeDBModel
	err := db.BunDB.NewSelect().
		Model(&transmissionTypes).
		Scan(ctx)

	return transmissionTypes, err
}

func (r *BunTransmissionTypeRepository) GetTransmissionTypeByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*TransmissionTypeDBModel, error) {
	transmissionType := new(TransmissionTypeDBModel)
	err := db.BunDB.NewSelect().
		Model(transmissionType).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return transmissionType, nil
}

func (r *BunTransmissionTypeRepository) CreateTransmissionTypeRepository(
	ctx context.Context,
	transmissionType *TransmissionTypeDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(transmissionType).
		Exec(ctx)
	return err
}

func (r *BunTransmissionTypeRepository) UpdateTransmissionTypeRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateTransmissionTypeRequest,
) (*TransmissionTypeDBModel, error) {
	transmissionType, err := r.GetTransmissionTypeByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		transmissionType.Code = *request.Code
	}
	_, err = db.BunDB.NewUpdate().
		Model(transmissionType).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return transmissionType, nil
}

func (r *BunTransmissionTypeRepository) GetTransmissionTypeByCodeRepository(
	ctx context.Context,
	code string,
) (*TransmissionTypeDBModel, error) {
	transmissionType := new(TransmissionTypeDBModel)
	err := db.BunDB.NewSelect().
		Model(transmissionType).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return transmissionType, nil
}
