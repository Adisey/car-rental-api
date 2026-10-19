package car_drive_types

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func GetDriveTypesService(ctx context.Context) ([]api_models.DriveType, error) {
	repo := NewDriveTypeRepository()
	dbItems, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]api_models.DriveType, 0, len(dbItems))
	for _, item := range dbItems {
		items = append(items, api_models.DriveType{
			Id:   item.ID.String(),
			Code: item.Code,
		})
	}
	return items, nil
}

func GetDriveTypeByIDService(
	ctx context.Context,
	id string,
) (*api_models.DriveType, error) {
	repo := NewDriveTypeRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetByIDRepository(ctx, itemID)
	if err != nil {
		return nil, err
	}
	return &api_models.DriveType{
		Id:   dbItem.ID.String(),
		Code: dbItem.Code,
	}, nil
}

func CreateDriveTypeService(
	ctx context.Context,
	request api_models.CreateDriveTypeRequest,
) (*api_models.DriveType, error) {
	validationResult := shared_validation.ValidateCreateCodeRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := NewDriveTypeRepository()
	existItem, errExist := repo.GetByCodeRepository(
		ctx,
		request.Code,
	)
	if errExist == nil {
		return &api_models.DriveType{
			Id:   existItem.ID.String(),
			Code: existItem.Code,
		}, nil
	}
	newItem := &DriveTypeDBModel{
		Code: request.Code,
	}
	errNew := repo.CreateRepository(ctx, newItem)
	if errNew != nil {
		return nil, errNew
	}
	return &api_models.DriveType{
		Id:   newItem.ID.String(),
		Code: newItem.Code,
	}, nil
}

func UpdateDriveTypeService(
	ctx context.Context,
	id string,
	request api_models.UpdateDriveTypeRequest,
) (*api_models.DriveType, error) {
	repo := NewDriveTypeRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	validationResult := shared_validation.ValidateUpdateCodeRequest(
		request,
	)
	if request.Code != nil {
		existItem, err := repo.GetByCodeRepository(
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
	itemUpdated, errUpdate := repo.UpdateRepository(
		ctx,
		itemID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}
	return &api_models.DriveType{
		Id:   itemUpdated.ID.String(),
		Code: itemUpdated.Code,
	}, nil
}
