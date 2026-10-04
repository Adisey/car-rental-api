package repositories

import (
	"context"

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