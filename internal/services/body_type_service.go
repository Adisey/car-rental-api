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

	dbBodyTypes, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	bodyTypes := make([]api_models.BodyType, 0, len(dbBodyTypes))

	for _, bodyType := range dbBodyTypes {
		bodyTypes = append(bodyTypes, api_models.BodyType{
			Id:   bodyType.ID.String(),
			Code: bodyType.Code,
		})
	}

	return bodyTypes, nil
}

func GetBodyTypeByIDService(
	ctx context.Context,
	id string,
) (*api_models.BodyType, error) {
	repo := repositories.NewBodyTypeRepository()

	bodyTypeID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbBodyType, err := repo.GetByIDRepository(ctx, bodyTypeID)
	if err != nil {
		return nil, err
	}

	return &api_models.BodyType{
		Id:   dbBodyType.ID.String(),
		Code: dbBodyType.Code,
	}, nil
}

func CreateBodyTypeService(
	ctx context.Context,
	request api_models.CreateBodyTypeRequest,
) (*api_models.BodyType, error) {

	validationResult := validation.ValidateCreateBodyTypeRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	repo := repositories.NewBodyTypeRepository()

	existingBodyType, errExist := repo.GetByCodeRepository(
		ctx,
		request.Code,
	)

	if errExist == nil {
		return &api_models.BodyType{
			Id:   existingBodyType.ID.String(),
			Code: existingBodyType.Code,
		}, nil
	}

	bodyType := &db_models.BodyType{
		Code: request.Code,
	}

	errNew := repo.CreateRepository(ctx, bodyType)
	if errNew != nil {
		return nil, errNew
	}

	return &api_models.BodyType{
		Id:   bodyType.ID.String(),
		Code: bodyType.Code,
	}, nil
}

func UpdateBodyTypeService(
	ctx context.Context,
	id string,
	request api_models.UpdateBodyTypeRequest,
) (*api_models.BodyType, error) {

	repo := repositories.NewBodyTypeRepository()

	bodyTypeID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}

	validationResult := validation.ValidateUpdateBodyTypeRequest(
		request,
	)

	if request.Code != nil {

		existingBodyType, err := repo.GetByCodeRepository(
			ctx,
			*request.Code,
		)

		if err == nil &&
			existingBodyType.ID != bodyTypeID {

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

	bodyType, errUpdate := repo.UpdateRepository(
		ctx,
		bodyTypeID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return &api_models.BodyType{
		Id:   bodyType.ID.String(),
		Code: bodyType.Code,
	}, nil
}
