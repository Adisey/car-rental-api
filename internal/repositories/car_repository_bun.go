package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BunCarRepository struct{}

func NewCarRepository() *BunCarRepository {
	return &BunCarRepository{}
}

func (r *BunCarRepository) GetAll(ctx context.Context) ([]db_models.Car, error) {
	var cars []db_models.Car

	err := db.BunDB.NewSelect().
		Model(&cars).
		Relation("Color").
		Where("deleted_at IS NULL").
		Scan(ctx)

	return cars, err
}

func (r *BunCarRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Car, error) {
	car := new(db_models.Car)
	log.Println("Repository GET /cars/", id)

	err := db.BunDB.NewSelect().
		Model(car).
		Relation("Color").
		Where("car.id = ?", id).
		Where("deleted_at IS NULL").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *BunCarRepository) CreateRepository(
	ctx context.Context,
	car *db_models.Car,
) error {

	_, err := db.BunDB.NewInsert().
		Model(car).
		Exec(ctx)

	return err
}

func (r *BunCarRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateCarRequest,
	colorID *uuid.UUID,
	shouldUpdateColor bool,
) (*db_models.Car, error) {
	car, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.Name != nil {
		car.Name = *request.Name
	}
	if request.Description != nil {
		car.Description = request.Description
	}
	if shouldUpdateColor {
		car.ColorID = colorID
	}
	car.UpdatedAt = time.Now().UTC()
	_, err = db.BunDB.NewUpdate().
		Model(car).
		WherePK().
		Where("deleted_at IS NULL").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (r *BunCarRepository) DeleteRepository(
	ctx context.Context,
	id uuid.UUID,
) error {

	car, err := r.GetByIDRepository(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now().UTC()

	car.DeletedAt = &now
	car.UpdatedAt = now

	_, err = db.BunDB.NewUpdate().
		Model(car).
		WherePK().
		Where("deleted_at IS NULL").
		Exec(ctx)

	return err
}
