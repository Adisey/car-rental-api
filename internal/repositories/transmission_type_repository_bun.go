package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BunTransmissionTypeRepository struct{}

func NewTransmissionTypeRepository() *BunTransmissionTypeRepository {
	return &BunTransmissionTypeRepository{}
}

func (r *BunTransmissionTypeRepository) GetAll(ctx context.Context) ([]db_models.TransmissionType, error) {
	var transmissionTypes []db_models.TransmissionType
	err := db.BunDB.NewSelect().
		Model(&transmissionTypes).
		Scan(ctx)

	return transmissionTypes, err
}

func (r *BunTransmissionTypeRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.TransmissionType, error) {
	transmissionType := new(db_models.TransmissionType)
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

func (r *BunTransmissionTypeRepository) CreateRepository(
	ctx context.Context,
	transmissionType *db_models.TransmissionType,
) error {
	_, err := db.BunDB.NewInsert().
		Model(transmissionType).
		Exec(ctx)
	return err
}

func (r *BunTransmissionTypeRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateTransmissionTypeRequest,
) (*db_models.TransmissionType, error) {
	transmissionType, err := r.GetByIDRepository(ctx, id)
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

func (r *BunTransmissionTypeRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*db_models.TransmissionType, error) {
	transmissionType := new(db_models.TransmissionType)
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
