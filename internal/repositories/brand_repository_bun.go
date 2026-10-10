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

type BunBrandRepository struct{}

func NewBrandRepository() *BunBrandRepository {
	return &BunBrandRepository{}
}

func (r *BunBrandRepository) GetAll(ctx context.Context) ([]db_models.Brand, error) {
	var Brands []db_models.Brand

	err := db.BunDB.NewSelect().
		Model(&Brands).
		Scan(ctx)

	return Brands, err
}

func (r *BunBrandRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Brand, error) {
	Brand := new(db_models.Brand)

	err := db.BunDB.NewSelect().
		Model(Brand).
		Where("id = ?", id).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return Brand, nil
}

func (r *BunBrandRepository) CreateRepository(
	ctx context.Context,
	Brand *db_models.Brand,
) error {

	_, err := db.BunDB.NewInsert().
		Model(Brand).
		Exec(ctx)

	return err
}

func (r *BunBrandRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateBrandRequest,
) (*db_models.Brand, error) {

	Brand, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Name != nil {
		Brand.Name = *request.Name
	}
	_, err = db.BunDB.NewUpdate().
		Model(Brand).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return Brand, nil
}

func (r *BunBrandRepository) GetByNameRepository(
	ctx context.Context,
	code string,
) (*db_models.Brand, error) {

	brand := new(db_models.Brand)

	err := db.BunDB.NewSelect().
		Model(brand).
		Where("name = ?", code).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return brand, nil
}
