package car_body

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	shared_validation "github.com/Adisey/car-rental-api/internal/shared/validation"
)

func GetBodyTypesService(ctx context.Context) ([]api_models.BodyType, error) {
	repo := NewBodyTypeRepository()
	dbItems, err := repo.GetBodyTypeAll(ctx)
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
	repo := NewBodyTypeRepository()
	itemID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbItem, err := repo.GetBodyTypeByIDRepository(ctx, itemID)
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
	validationResult := shared_validation.ValidateCreateCodeRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	repo := NewBodyTypeRepository()
	existItem, errExist := repo.GetBodyTypeByCodeRepository(
		ctx,
		request.Code,
	)
	if errExist == nil {
		return &api_models.BodyType{
			Id:   existItem.ID.String(),
			Code: existItem.Code,
		}, nil
	}
	newItem := &BodyTypeDBModel{
		Code: request.Code,
	}
	errNew := repo.CreateBodyTypeRepository(ctx, newItem)
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
	repo := NewBodyTypeRepository()
	itemID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}
	validationResult := shared_validation.ValidateUpdateCodeRequest(
		request,
	)
	if request.Code != nil {
		existItem, err := repo.GetBodyTypeByCodeRepository(
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
	itemUpdated, errUpdate := repo.UpdateBodyTypeRepository(
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
