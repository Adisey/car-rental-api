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
	var Colors []ColorDBModel

	err := db.BunDB.NewSelect().
		Model(&Colors).
		Scan(ctx)

	return Colors, err
}

func (r *BunColorRepository) GetColorByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*ColorDBModel, error) {
	Color := new(ColorDBModel)
	err := db.BunDB.NewSelect().
		Model(Color).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return Color, nil
}

func (r *BunColorRepository) CreateColorRepository(
	ctx context.Context,
	Color *ColorDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(Color).
		Exec(ctx)

	return err
}

func (r *BunColorRepository) UpdateColorRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateColorRequest,
) (*ColorDBModel, error) {
	Color, err := r.GetColorByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		Color.Code = *request.Code
		Color.Hex = request.Hex
	}
	_, err = db.BunDB.NewUpdate().
		Model(Color).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return Color, nil
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
