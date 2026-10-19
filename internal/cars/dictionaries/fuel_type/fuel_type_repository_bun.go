package car_fuel

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

type BunFuelTypeRepository struct{}

func NewFuelTypeRepository() *BunFuelTypeRepository {
	return &BunFuelTypeRepository{}
}

func (r *BunFuelTypeRepository) GetAll(ctx context.Context) ([]FuelTypeDBModel, error) {
	var fuelTypes []FuelTypeDBModel
	err := db.BunDB.NewSelect().
		Model(&fuelTypes).
		Scan(ctx)

	return fuelTypes, err
}

func (r *BunFuelTypeRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*FuelTypeDBModel, error) {
	fuelType := new(FuelTypeDBModel)
	err := db.BunDB.NewSelect().
		Model(fuelType).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return fuelType, nil
}

func (r *BunFuelTypeRepository) CreateRepository(
	ctx context.Context,
	fuelType *FuelTypeDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(fuelType).
		Exec(ctx)
	return err
}

func (r *BunFuelTypeRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateFuelTypeRequest,
) (*FuelTypeDBModel, error) {
	fuelType, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		fuelType.Code = *request.Code
	}
	_, err = db.BunDB.NewUpdate().
		Model(fuelType).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return fuelType, nil
}

func (r *BunFuelTypeRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*FuelTypeDBModel, error) {
	fuelType := new(FuelTypeDBModel)
	err := db.BunDB.NewSelect().
		Model(fuelType).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return fuelType, nil
}
