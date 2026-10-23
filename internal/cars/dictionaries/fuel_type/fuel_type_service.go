package car_fuel

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func GetFuelTypesService(ctx context.Context) ([]api_models.FuelType, error) {
	repo := NewFuelTypeRepository()
	dbItems, err := repo.GetFuelTypeAllRepository(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]api_models.FuelType, 0, len(dbItems))
	for _, item := range dbItems {
		items = append(items, api_models.FuelType{
			Id:   item.ID.String(),
			Code: item.Code,
		})
	}
	return items, nil
}

func GetFuelTypeByIDService(
	ctx context.Context,
	id string,
) (*api_models.FuelType, error) {
	repo := NewFuelTypeRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetFuelTypeByIDRepository(ctx, itemID)
	if err != nil {
		return nil, err
	}
	return &api_models.FuelType{
		Id:   dbItem.ID.String(),
		Code: dbItem.Code,
	}, nil
}

func CreateFuelTypeService(
	ctx context.Context,
	request api_models.CreateCodeEntityRequest,
) (*api_models.FuelType, error) {
	validationResult := shared_validation.ValidateCreateCodeRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := NewFuelTypeRepository()
	existItem, errExist := repo.GetFuelTypeByCodeRepository(
		ctx,
		request.Code,
	)
	if errExist == nil {
		return &api_models.FuelType{
			Id:   existItem.ID.String(),
			Code: existItem.Code,
		}, nil
	}
	newItem := &FuelTypeDBModel{
		Code: request.Code,
	}
	errNew := repo.CreateFuelTypeRepository(ctx, newItem)
	if errNew != nil {
		return nil, errNew
	}
	return &api_models.FuelType{
		Id:   newItem.ID.String(),
		Code: newItem.Code,
	}, nil
}

func UpdateFuelTypeService(
	ctx context.Context,
	id string,
	request api_models.UpdateCodeEntityRequest,
) (*api_models.FuelType, error) {
	repo := NewFuelTypeRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	validationResult := shared_validation.ValidateUpdateCodeRequest(
		request,
	)
	if request.Code != nil {
		existItem, err := repo.GetFuelTypeByCodeRepository(
			ctx,
			*request.Code,
		)
		if err == nil &&
			existItem.ID != itemID {

			validationResult.Add(
				"code",
				"already_exists",
				nil,
			)
		}
	}
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	itemUpdated, errUpdate := repo.UpdateFuelTypeRepository(
		ctx,
		itemID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}
	return &api_models.FuelType{
		Id:   itemUpdated.ID.String(),
		Code: itemUpdated.Code,
	}, nil
}
