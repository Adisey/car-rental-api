package customers

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

type CustomerRepository interface {
	GetCustomersAllRepository(ctx context.Context) ([]CustomerDBModel, error)

	GetCustomerByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*CustomerDBModel, error)

	GetCustomerByNameRepository(
		ctx context.Context,
		name string,
	) (*CustomerDBModel, error)

	CreateCustomerRepository(
		ctx context.Context,
		Customer *CustomerDBModel,
	) (*CustomerDBModel, error)

	UpdateCustomerRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateCustomerRequest,
	) (*CustomerDBModel, error)
}
