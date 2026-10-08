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

type BunColorRepository struct{}

func NewColorRepository() *BunColorRepository {
	return &BunColorRepository{}
}

func (r *BunColorRepository) GetAll(ctx context.Context) ([]db_models.Color, error) {
	var Colors []db_models.Color

	err := db.BunDB.NewSelect().
		Model(&Colors).
		Scan(ctx)

	return Colors, err
}

func (r *BunColorRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Color, error) {
	Color := new(db_models.Color)

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

func (r *BunColorRepository) CreateRepository(
	ctx context.Context,
	Color *db_models.Color,
) error {

	_, err := db.BunDB.NewInsert().
		Model(Color).
		Exec(ctx)

	return err
}

func (r *BunColorRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateColorRequest,
) (*db_models.Color, error) {

	Color, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Code != nil {
		Color.Code = *request.Code
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

func (r *BunColorRepository) GetByCodeRepository(
	ctx context.Context,
	code string,
) (*db_models.Color, error) {

	color := new(db_models.Color)

	err := db.BunDB.NewSelect().
		Model(color).
		Where("code = ?", code).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return color, nil
}
