package car_color

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
)

func GetColorsService(ctx context.Context) ([]api_models.Color, error) {
	repo := NewColorRepository()

	dbColors, err := repo.GetColorAll(ctx)
	if err != nil {
		return nil, err
	}

	colors := make([]api_models.Color, 0, len(dbColors))

	for _, color := range dbColors {
		colors = append(colors, api_models.Color{
			Id:   color.ID.String(),
			Code: color.Code,
			Hex:  color.Hex,
		})
	}

	return colors, nil
}

func GetColorByIDService(
	ctx context.Context,
	id string,
) (*api_models.Color, error) {
	repo := NewColorRepository()

	colorID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbColor, err := repo.GetColorByIDRepository(ctx, colorID)
	if err != nil {
		return nil, err
	}
	return &api_models.Color{
		Id:   dbColor.ID.String(),
		Code: dbColor.Code,
		Hex:  dbColor.Hex,
	}, nil
}

func CreateColorService(
	ctx context.Context,
	request api_models.CreateColorRequest,
) (*api_models.Color, error) {

	validationResult := ValidateCreateColorRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	repo := NewColorRepository()

	existingColor, errExist := repo.GetColorByCodeRepository(
		ctx,
		request.Code,
	)

	if errExist == nil {
		return &api_models.Color{
			Id:   existingColor.ID.String(),
			Code: existingColor.Code,
			Hex:  existingColor.Hex,
		}, nil
	}

	color := &ColorDBModel{
		Code: request.Code,
		Hex:  request.Hex,
	}

	errNew := repo.CreateColorRepository(ctx, color)
	if errNew != nil {
		return nil, errNew
	}

	return &api_models.Color{
		Id:   color.ID.String(),
		Code: color.Code,
		Hex:  color.Hex,
	}, nil
}

func UpdateColorService(
	ctx context.Context,
	id string,
	request api_models.UpdateColorRequest,
) (*api_models.Color, error) {

	repo := NewColorRepository()

	colorID, errId := uuid.Parse(id)
	if errId != nil {
		return nil, errId
	}

	validationResult := ValidateUpdateColorRequest(
		request,
	)

	if request.Code != nil {

		existingColor, err := repo.GetColorByCodeRepository(
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

	color, errUpdate := repo.UpdateColorRepository(
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
		Hex:  color.Hex,
	}, nil
}
