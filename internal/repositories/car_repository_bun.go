package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"

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
		Where("id = ?", id).
		Scan(ctx)

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
