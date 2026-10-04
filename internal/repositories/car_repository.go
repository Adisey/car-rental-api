package repositories

import (
	"context"

	"github.com/Adisey/car-rental-api/internal/db_models"
)

type CarRepository interface {
	GetAll(ctx context.Context) ([]db_models.Car, error)
}
