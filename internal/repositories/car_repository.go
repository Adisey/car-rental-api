package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/db_models"
)

type CarRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.Car, error)
	GetByIDRepository(ctx context.Context, id uuid.UUID) (*db_models.Car, error)
	CreateRepository(ctx context.Context, car *db_models.Car) error
}
