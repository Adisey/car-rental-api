package car_model

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/google/uuid"
)

type BunModelRepository struct{}

func NewModelRepository() *BunModelRepository {
	return &BunModelRepository{}
}

func (r *BunModelRepository) GetModelAllRepository(
	ctx context.Context,
	brandID *uuid.UUID,
) ([]ModelDBModel, error) {
	var models []ModelDBModel
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

func (r *BunModelRepository) GetModelByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*ModelDBModel, error) {
	model := new(ModelDBModel)
	err := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("models.id = ?", id).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *BunModelRepository) GetModelByNameRepository(
	ctx context.Context,
	name string,
	brandID *uuid.UUID,
) (*ModelDBModel, error) {
	model := new(ModelDBModel)
	query := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("LOWER(models.name) = LOWER(?)", name)
	if brandID != nil {
		query = query.Where("models.brand_id = ?", brandID)
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
	model *ModelDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(model).
		Exec(ctx)
	return err
}

func (r *BunModelRepository) UpdateModelRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateModelRequest,
	brandID *uuid.UUID,
	shouldUpdateBrand bool,
) (*ModelDBModel, error) {
	model, err := r.GetModelByIDRepository(
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
