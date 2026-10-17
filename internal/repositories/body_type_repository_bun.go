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
	var bodyTypes []db_models.BodyType
	err := db.BunDB.NewSelect().
		Model(&bodyTypes).
		Scan(ctx)

	return bodyTypes, err
}

func (r *BunBodyTypeRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.BodyType, error) {
	bodyType := new(db_models.BodyType)
	err := db.BunDB.NewSelect().
		Model(bodyType).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return bodyType, nil
}

func (r *BunBodyTypeRepository) CreateRepository(
	ctx context.Context,
	bodyType *db_models.BodyType,
) error {
	_, err := db.BunDB.NewInsert().
		Model(bodyType).
		Exec(ctx)
	return err
}

func (r *BunBodyTypeRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateBodyTypeRequest,
) (*db_models.BodyType, error) {
	bodyType, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		bodyType.Code = *request.Code
	}
	_, err = db.BunDB.NewUpdate().
		Model(bodyType).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return bodyType, nil
}

func (r *BunBodyTypeRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*db_models.BodyType, error) {
	bodyType := new(db_models.BodyType)
	err := db.BunDB.NewSelect().
		Model(bodyType).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return bodyType, nil
}
