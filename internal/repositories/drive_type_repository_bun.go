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

type BunDriveTypeRepository struct{}

func NewDriveTypeRepository() *BunDriveTypeRepository {
	return &BunDriveTypeRepository{}
}

func (r *BunDriveTypeRepository) GetAll(ctx context.Context) ([]db_models.DriveType, error) {
	var driveType []db_models.DriveType
	err := db.BunDB.NewSelect().
		Model(&driveType).
		Scan(ctx)

	return driveType, err
}

func (r *BunDriveTypeRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.DriveType, error) {
	driveType := new(db_models.DriveType)
	err := db.BunDB.NewSelect().
		Model(driveType).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return driveType, nil
}

func (r *BunDriveTypeRepository) CreateRepository(
	ctx context.Context,
	driveType *db_models.DriveType,
) error {
	_, err := db.BunDB.NewInsert().
		Model(driveType).
		Exec(ctx)
	return err
}

func (r *BunDriveTypeRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateDriveTypeRequest,
) (*db_models.DriveType, error) {
	driveType, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		driveType.Code = *request.Code
	}
	_, err = db.BunDB.NewUpdate().
		Model(driveType).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return driveType, nil
}

func (r *BunDriveTypeRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*db_models.DriveType, error) {
	driveType := new(db_models.DriveType)
	err := db.BunDB.NewSelect().
		Model(driveType).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return driveType, nil
}
