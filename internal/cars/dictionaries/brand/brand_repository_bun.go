package car_brand

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

type BunBrandRepository struct{}

func NewBrandRepository() *BunBrandRepository {
	return &BunBrandRepository{}
}

func (r *BunBrandRepository) GetBrandAll(ctx context.Context) ([]BrandDBModel, error) {
	var Brands []BrandDBModel

	err := db.BunDB.NewSelect().
		Model(&Brands).
		Scan(ctx)

	return Brands, err
}

func (r *BunBrandRepository) GetBrandByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*BrandDBModel, error) {
	Brand := new(BrandDBModel)

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

func (r *BunBrandRepository) CreateBrandRepository(
	ctx context.Context,
	Brand *BrandDBModel,
) error {

	_, err := db.BunDB.NewInsert().
		Model(Brand).
		Exec(ctx)

	return err
}

func (r *BunBrandRepository) UpdateBrandRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateBrandRequest,
) (*BrandDBModel, error) {

	Brand, err := r.GetBrandByIDRepository(ctx, id)
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

func (r *BunBrandRepository) GetBrandByNameRepository(
	ctx context.Context,
	name string,
) (*BrandDBModel, error) {

	brand := new(BrandDBModel)

	err := db.BunDB.NewSelect().
		Model(brand).
		Where("LOWER(name) = LOWER(?)", name).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return brand, nil
}
