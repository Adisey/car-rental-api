package customers

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func mapCustomerApiFields(item *CustomerDBModel) api_models.Customer {
	return api_models.Customer{
		Id:        item.ID.String(),
		FirstName: item.FirstName,
		LastName:  item.LastName,
		Phone:     item.Phone,
		Email:     item.Email,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func mapCreateCustomerRequest(
	request api_models.CreateCustomerRequest,
) *CustomerDBModel {
	return &CustomerDBModel{
		FirstName: strings.TrimSpace(request.FirstName),
		LastName:  request.LastName,
		Phone:     request.Phone,
		Email:     request.Email,
	}
}

func applyUpdateCustomerRequest(
	customer *CustomerDBModel,
	request api_models.UpdateCustomerRequest,
) {
	if request.FirstName != nil {
		customer.FirstName = strings.TrimSpace(*request.FirstName)
	}
	if request.LastName != nil {
		customer.LastName = request.LastName
	}
	if request.Phone != nil {
		customer.Phone = request.Phone
	}
	if request.Email != nil {
		customer.Email = request.Email
	}
}

func GetCustomersService(ctx context.Context) ([]api_models.Customer, error) {
	repo := NewCustomerRepository()
	dbItems, err := repo.GetCustomerAll(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]api_models.Customer, 0, len(dbItems))
	for _, item := range dbItems {
		customer := mapCustomerApiFields(&item)
		items = append(items, customer)
	}
	return items, nil
}

func GetCustomerByIDService(
	ctx context.Context,
	id string,
) (*api_models.Customer, error) {
	repo := NewCustomerRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetCustomerByIDRepository(ctx, itemID)
	if err != nil {
		return nil, err
	}
	customer := mapCustomerApiFields(dbItem)
	return &customer, nil
}

func FindExistingCreateCustomer(
	ctx context.Context,
	request api_models.CreateCustomerRequest,
) (*api_models.Customer, error) {
	repo := NewCustomerRepository()
	if request.Phone != nil && *request.Phone != "" {
		customer, err := repo.GetCustomerByPhoneRepository(ctx, *request.Phone)
		if err == nil {
			apiCustomer := mapCustomerApiFields(customer)
			return &apiCustomer, api_errors.ErrCustomerExist
		}
	}
	if request.Email != nil && *request.Email != "" {
		customer, err := repo.GetCustomerByEmailRepository(ctx, *request.Email)
		if err == nil {
			apiCustomer := mapCustomerApiFields(customer)
			return &apiCustomer, api_errors.ErrCustomerExist
		}
	}
	customer, err := repo.GetCustomerByNameRepository(ctx, request.FirstName, request.LastName)
	if err == nil {
		apiCustomer := mapCustomerApiFields(customer)
		return &apiCustomer, api_errors.ErrCustomerExist
	}
	return nil, nil
}

func CreateCustomerService(
	ctx context.Context,
	request api_models.CreateCustomerRequest,
) (*api_models.Customer, error) {
	validationResult := ValidateCreateCustomerRequest(request)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := NewCustomerRepository()
	existItem, errExist := FindExistingCreateCustomer(ctx, request)
	if errExist != nil {
		return existItem, errExist
	}
	newItem := mapCreateCustomerRequest(request)
	errNew := repo.CreateCustomerRepository(ctx, newItem)
	if errNew != nil {
		return nil, errNew
	}
	dbCustomer, errCreate := repo.GetCustomerByIDRepository(ctx, newItem.ID)
	customer := mapCustomerApiFields(dbCustomer)
	return &customer, errCreate
}

func FindExistingUpdateCustomer(
	ctx context.Context,
	id string,
	request api_models.UpdateCustomerRequest,
) (*shared_validation.ValidationResult, error) {
	result := shared_validation.NewValidationResult()
	repo := NewCustomerRepository()
	currentID, _ := uuid.Parse(id)
	currentCustomer, errCurrent := repo.GetCustomerByIDRepository(ctx, currentID)
	if errCurrent != nil {
		return nil, errCurrent
	}
	if request.Phone != nil && *request.Phone != "" {
		customer, err := repo.GetCustomerByPhoneRepository(ctx, *request.Phone)
		if err == nil && customer.ID.String() != id {
			result.Add("phone", "already_exist", nil)
		}
	}
	if request.Email != nil && *request.Email != "" {
		customer, err := repo.GetCustomerByEmailRepository(ctx, *request.Email)
		if err == nil && customer.ID.String() != id {
			result.Add("email", "already_exist", nil)
		}
	}
	firstName := currentCustomer.FirstName
	if request.FirstName != nil {
		firstName = *request.FirstName
	}
	lastName := currentCustomer.LastName
	if request.LastName != nil {
		lastName = request.LastName
	}
	customer, err := repo.GetCustomerByNameRepository(ctx, firstName, lastName)
	if err == nil && customer.ID.String() != id {
		result.Add("first_name", "already_exist", nil)
		result.Add("last_name", "already_exist", nil)
	}
	return result, nil
}

func UpdateCustomerService(
	ctx context.Context,
	id string,
	request api_models.UpdateCustomerRequest,
) (*api_models.Customer, error) {
	repo := NewCustomerRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	rulesValidationResult := ValidateUpdateCustomerRequest(request)
	dbValidationResult, dbValidationError := FindExistingUpdateCustomer(ctx, id, request)
	if dbValidationError != nil {
		return nil, dbValidationError
	}
	validationResult := shared_validation.NewValidationResult()
	validationResult.Merge(rulesValidationResult)
	validationResult.Merge(dbValidationResult)
	if validationResult.HasErrors() {
		return nil, rulesValidationResult
	}
	itemUpdated, errUpdate := repo.UpdateCustomerRepository(
		ctx,
		itemID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}
	customer := mapCustomerApiFields(itemUpdated)
	return &customer, nil
}
