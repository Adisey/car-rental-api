package car_transmission

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func GetTransmissionTypesService(ctx context.Context) ([]api_models.TransmissionType, error) {
	repo := NewTransmissionTypeRepository()
	dbItems, err := repo.GetTransmissionTypeAllRepository(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]api_models.TransmissionType, 0, len(dbItems))
	for _, item := range dbItems {
		items = append(items, api_models.TransmissionType{
			Id:   item.ID.String(),
			Code: item.Code,
		})
	}
	return items, nil
}

func GetTransmissionTypeByIDService(
	ctx context.Context,
	id string,
) (*api_models.TransmissionType, error) {
	repo := NewTransmissionTypeRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetTransmissionTypeByIDRepository(ctx, itemID)
	if err != nil {
		return nil, err
	}
	return &api_models.TransmissionType{
		Id:   dbItem.ID.String(),
		Code: dbItem.Code,
	}, nil
}

func CreateTransmissionTypeService(
	ctx context.Context,
	request api_models.CreateCodeEntityRequest,
) (*api_models.TransmissionType, error) {
	validationResult := shared_validation.ValidateCreateCodeRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := NewTransmissionTypeRepository()
	existItem, errExist := repo.GetTransmissionTypeByCodeRepository(
		ctx,
		request.Code,
	)
	if errExist == nil {
		return &api_models.TransmissionType{
			Id:   existItem.ID.String(),
			Code: existItem.Code,
		}, nil
	}
	newItem := &TransmissionTypeDBModel{
		Code: request.Code,
	}
	errNew := repo.CreateTransmissionTypeRepository(ctx, newItem)
	if errNew != nil {
		return nil, errNew
	}
	return &api_models.TransmissionType{
		Id:   newItem.ID.String(),
		Code: newItem.Code,
	}, nil
}

func UpdateTransmissionTypeService(
	ctx context.Context,
	id string,
	request api_models.UpdateCodeEntityRequest,
) (*api_models.TransmissionType, error) {
	repo := NewTransmissionTypeRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	validationResult := shared_validation.ValidateUpdateCodeRequest(
		request,
	)
	if request.Code != nil {
		existItem, err := repo.GetTransmissionTypeByCodeRepository(
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
	itemUpdated, errUpdate := repo.UpdateTransmissionTypeRepository(
		ctx,
		itemID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}
	return &api_models.TransmissionType{
		Id:   itemUpdated.ID.String(),
		Code: itemUpdated.Code,
	}, nil
}
