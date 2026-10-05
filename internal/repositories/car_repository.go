package repositories

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/db_models"
)

type CarRepository interface {
	GetAll(ctx context.Context) ([]db_models.Car, error)
	GetByID(ctx context.Context, id uuid.UUID) (*db_models.Car, error)
}
