package car_color

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

type BunColorRepository struct{}

func NewColorRepository() *BunColorRepository {
	return &BunColorRepository{}
}

func (r *BunColorRepository) GetColorAll(ctx context.Context) ([]ColorDBModel, error) {
	var colors []ColorDBModel

	err := db.BunDB.NewSelect().
		Model(&colors).
		Scan(ctx)

	return colors, err
}

func (r *BunColorRepository) GetColorByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*ColorDBModel, error) {
	color := new(ColorDBModel)
	err := db.BunDB.NewSelect().
		Model(color).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return color, nil
}

func (r *BunColorRepository) CreateColorRepository(
	ctx context.Context,
	color *ColorDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(color).
		Exec(ctx)

	return err
}

func (r *BunColorRepository) UpdateColorRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateColorRequest,
) (*ColorDBModel, error) {
	color, err := r.GetColorByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		color.Code = *request.Code
		color.Hex = request.Hex
	}
	_, err = db.BunDB.NewUpdate().
		Model(color).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return color, nil
}

func (r *BunColorRepository) GetColorByCodeRepository(
	ctx context.Context,
	code string,
) (*ColorDBModel, error) {
	color := new(ColorDBModel)
	err := db.BunDB.NewSelect().
		Model(color).
		Where("LOWER(code) = LOWER(?)", code).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return color, nil
}
