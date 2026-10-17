package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetBodyTypesService(ctx context.Context) ([]api_models.BodyType, error) {
	repo := repositories.NewBodyTypeRepository()
	dbItems, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]api_models.BodyType, 0, len(dbItems))
	for _, item := range dbItems {
		items = append(items, api_models.BodyType{
			Id:   item.ID.String(),
			Code: item.Code,
		})
	}
	return items, nil
}

func GetBodyTypeByIDService(
	ctx context.Context,
	id string,
) (*api_models.BodyType, error) {
	repo := repositories.NewBodyTypeRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetByIDRepository(ctx, itemID)
	if err != nil {
		return nil, err
	}
	return &api_models.BodyType{
		Id:   dbItem.ID.String(),
		Code: dbItem.Code,
	}, nil
}

func CreateBodyTypeService(
	ctx context.Context,
	request api_models.CreateCodeEntityRequest,
) (*api_models.BodyType, error) {
	validationResult := validation.ValidateCreateCodeRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := repositories.NewBodyTypeRepository()
	existItem, errExist := repo.GetByCodeRepository(
		ctx,
		request.Code,
	)
	if errExist == nil {
		return &api_models.BodyType{
			Id:   existItem.ID.String(),
			Code: existItem.Code,
		}, nil
	}
	newItem := &db_models.BodyType{
		Code: request.Code,
	}
	errNew := repo.CreateRepository(ctx, newItem)
	if errNew != nil {
		return nil, errNew
	}
	return &api_models.BodyType{
		Id:   newItem.ID.String(),
		Code: newItem.Code,
	}, nil
}

func UpdateBodyTypeService(
	ctx context.Context,
	id string,
	request api_models.UpdateCodeEntityRequest,
) (*api_models.BodyType, error) {
	repo := repositories.NewBodyTypeRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	validationResult := validation.ValidateUpdateCodeRequest(
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
	return &api_models.BodyType{
		Id:   itemUpdated.ID.String(),
		Code: itemUpdated.Code,
	}, nil
}
