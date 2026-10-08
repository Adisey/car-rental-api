package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetColorsService(ctx context.Context) ([]api_models.Color, error) {
	repo := repositories.NewColorRepository()

	dbColors, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	colors := make([]api_models.Color, 0, len(dbColors))

	for _, color := range dbColors {
		colors = append(colors, api_models.Color{
			Id:   color.ID.String(),
			Code: color.Code,
		})
	}

	return colors, nil
}

func GetColorByIDService(
	ctx context.Context,
	id string,
) (*api_models.Color, error) {
	repo := repositories.NewColorRepository()

	colorID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbColor, err := repo.GetByIDRepository(ctx, colorID)
	if err != nil {
		return nil, err
	}

	return &api_models.Color{
		Id:   dbColor.ID.String(),
		Code: dbColor.Code,
	}, nil
}

func CreateColorService(
	ctx context.Context,
	request api_models.CreateColorRequest,
) (*api_models.Color, error) {

	validationResult := validation.ValidateCreateColorRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	repo := repositories.NewColorRepository()

	existingColor, errExist := repo.GetByCodeRepository(
		ctx,
		request.Code,
	)

	if errExist == nil {
		return &api_models.Color{
			Id:   existingColor.ID.String(),
			Code: existingColor.Code,
		}, nil
	}

	color := &db_models.Color{
		Code: request.Code,
	}

	errNew := repo.CreateRepository(ctx, color)
	if errNew != nil {
		return nil, errNew
	}

	return &api_models.Color{
		Id:   color.ID.String(),
		Code: color.Code,
	}, nil
}

func UpdateColorService(
	ctx context.Context,
	id string,
	request api_models.UpdateColorRequest,
) (*api_models.Color, error) {

	repo := repositories.NewColorRepository()

	colorID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}

	validationResult := validation.ValidateUpdateColorRequest(
		request,
	)

	if request.Code != nil {

		existingColor, err := repo.GetByCodeRepository(
			ctx,
			*request.Code,
		)

		if err == nil &&
			existingColor.ID != colorID {

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

	color, errUpdate := repo.UpdateRepository(
		ctx,
		colorID,
		request,
	)
	if errUpdate != nil {
		return nil, errUpdate
	}

	return &api_models.Color{
		Id:   color.ID.String(),
		Code: color.Code,
	}, nil
}
