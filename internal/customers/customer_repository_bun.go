package customers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db"
)

type BunCustomerRepository struct{}

func NewCustomerRepository() *BunCustomerRepository {
	return &BunCustomerRepository{}
}

func (r *BunCustomerRepository) GetCustomerAll(ctx context.Context) ([]CustomerDBModel, error) {
	var customers []CustomerDBModel
	err := db.BunDB.NewSelect().
		Model(&customers).
		Scan(ctx)
	return customers, err
}

func (r *BunCustomerRepository) GetCustomerByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*CustomerDBModel, error) {
	customer := new(CustomerDBModel)
	err := db.BunDB.NewSelect().
		Model(customer).
		Where("id = ?", id).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *BunCustomerRepository) CreateCustomerRepository(
	ctx context.Context,
	customer *CustomerDBModel,
) error {
	_, err := db.BunDB.NewInsert().
		Model(customer).
		Exec(ctx)
	return err
}

func (r *BunCustomerRepository) UpdateCustomerRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateCustomerRequest,
) (*CustomerDBModel, error) {
	customer, err := r.GetCustomerByIDRepository(ctx, id)
	if err != nil {
		return nil, err
	}
	if request.FirstName != nil {
		customer.FirstName = *request.FirstName
	}
	_, err = db.BunDB.NewUpdate().
		Model(customer).
		WherePK().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *BunCustomerRepository) GetCustomerByPhoneRepository(
	ctx context.Context,
	phone string,
) (*CustomerDBModel, error) {
	customer := new(CustomerDBModel)
	err := db.BunDB.NewSelect().
		Model(customer).
		Where("phone = ?", phone).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *BunCustomerRepository) GetCustomerByEmailRepository(
	ctx context.Context,
	email string,
) (*CustomerDBModel, error) {
	customer := new(CustomerDBModel)
	err := db.BunDB.NewSelect().
		Model(customer).
		Where("email = ?", email).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (r *BunCustomerRepository) GetCustomerByNameRepository(
	ctx context.Context,
	firstName string,
	lastName *string,
) (*CustomerDBModel, error) {
	customer := new(CustomerDBModel)
	query := db.BunDB.NewSelect().
		Model(customer).
		Where("first_name = ?", firstName)
	if lastName == nil {
		query = query.Where("last_name IS NULL")
	} else {
		query = query.Where(
			"last_name = ?",
			*lastName,
		)
	}
	err := query.Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return customer, nil
}
