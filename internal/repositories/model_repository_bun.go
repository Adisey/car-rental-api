package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/google/uuid"
)

type BunModelRepository struct{}

func NewModelRepository() *BunModelRepository {
	return &BunModelRepository{}
}

func (r *BunModelRepository) GetAllRepository(
	ctx context.Context,
	brandID *uuid.UUID,
) ([]db_models.Model, error) {

	var models []db_models.Model

	query := db.BunDB.NewSelect().
		Model(&models).
		Relation("Brand")

	if brandID != nil {
		query = query.Where(
			"model.brand_id = ?",
			*brandID,
		)
	}

	err := query.Scan(ctx)

	return models, err
}

func (r *BunModelRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Model, error) {

	model := new(db_models.Model)

	err := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("model.id = ?", id).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *BunModelRepository) GetByNameRepository(
	ctx context.Context,
	name string,
	brandID *uuid.UUID,
) (*db_models.Model, error) {

	model := new(db_models.Model)

	query := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("LOWER(model.name) = LOWER(?)", name)

	if brandID != nil {
		query = query.Where("model.brand_id = ?", brandID)
	}

	err := query.Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *BunModelRepository) CreateRepository(
	ctx context.Context,
	model *db_models.Model,
) error {

	_, err := db.BunDB.NewInsert().
		Model(model).
		Exec(ctx)

	return err
}

func (r *BunModelRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateModelRequest,
	brandID *uuid.UUID,
	shouldUpdateBrand bool,
) (*db_models.Model, error) {

	model, err := r.GetByIDRepository(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if request.Name != nil {
		model.Name = *request.Name
	}

	if shouldUpdateBrand {
		model.BrandID = brandID
	}

	_, err = db.BunDB.NewUpdate().
		Model(model).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return model, nil
}
