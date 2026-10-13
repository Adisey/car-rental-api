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

type BunBodyTypeRepository struct{}

func NewBodyTypeRepository() *BunBodyTypeRepository {
	return &BunBodyTypeRepository{}
}

func (r *BunBodyTypeRepository) GetAll(ctx context.Context) ([]db_models.BodyType, error) {
	var BodyTypes []db_models.BodyType
	err := db.BunDB.NewSelect().
		Model(&BodyTypes).
		Scan(ctx)

	return BodyTypes, err
}

func (r *BunBodyTypeRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.BodyType, error) {
	BodyType := new(db_models.BodyType)
	err := db.BunDB.NewSelect().
		Model(BodyType).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return BodyType, nil
}

func (r *BunBodyTypeRepository) CreateRepository(
	ctx context.Context,
	BodyType *db_models.BodyType,
) error {
	_, err := db.BunDB.NewInsert().
		Model(BodyType).
		Exec(ctx)
	return err
}

func (r *BunBodyTypeRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateBodyTypeRequest,
) (*db_models.BodyType, error) {
	BodyType, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		BodyType.Code = *request.Code
	}
	_, err = db.BunDB.NewUpdate().
		Model(BodyType).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return BodyType, nil
}

func (r *BunBodyTypeRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*db_models.BodyType, error) {
	brand := new(db_models.BodyType)
	err := db.BunDB.NewSelect().
		Model(brand).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return brand, nil
}
