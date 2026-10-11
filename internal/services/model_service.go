package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetModelsService(
	ctx context.Context,
	brandID *uuid.UUID,
) ([]api_models.Model, error) {

	repo := repositories.NewModelRepository()

	dbModels, err := repo.GetAllRepository(
		ctx,
		brandID,
	)

	if err != nil {
		return nil, err
	}

	models := make(
		[]api_models.Model,
		0,
		len(dbModels),
	)

	for _, model := range dbModels {

		var brand *api_models.NameReference

		if model.Brand != nil {
			brandID := model.Brand.ID.String()
			brandName := model.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}

		models = append(
			models,
			api_models.Model{
				Id:    model.ID.String(),
				Name:  model.Name,
				Brand: brand,
			},
		)
	}

	return models, nil
}

func GetModelByIDService(
	ctx context.Context,
	id string,
) (*api_models.Model, error) {

	modelID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	repo := repositories.NewModelRepository()

	dbModel, err := repo.GetByIDRepository(
		ctx,
		modelID,
	)

	if err != nil {
		return nil, err
	}

	var brand *api_models.NameReference

	if dbModel.Brand != nil {
		brandID := dbModel.Brand.ID.String()
		brandName := dbModel.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}

	return &api_models.Model{
		Id:    dbModel.ID.String(),
		Name:  dbModel.Name,
		Brand: brand,
	}, nil
}

func CreateModelService(
	ctx context.Context,
	request api_models.CreateModelRequest,
) (*api_models.Model, error) {

	validationResult := validation.ValidateCreateModelRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	var brandID *uuid.UUID

	if request.Brand != nil &&
		!IsEmptyNameReference(request.Brand) {

		var err error

		brandID, err = ResolveBrandID(
			ctx,
			request.Brand,
		)

		if err != nil {
			return nil, err
		}
	}

	repo := repositories.NewModelRepository()

	existingModel, err := repo.GetByNameRepository(
		ctx,
		request.Name,
		brandID,
	)
	if err == nil {
		var brand *api_models.NameReference
		if existingModel.Brand != nil {
			brandID := existingModel.Brand.ID.String()
			brandName := existingModel.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}
		return &api_models.Model{
			Id:    existingModel.ID.String(),
			Name:  existingModel.Name,
			Brand: brand,
		}, nil
	}

	model := &db_models.Model{
		Name:    request.Name,
		BrandID: brandID,
	}

	err = repo.CreateRepository(
		ctx,
		model,
	)

	if err != nil {
		return nil, err
	}

	model, err = repo.GetByIDRepository(
		ctx,
		model.ID,
	)

	if err != nil {
		return nil, err
	}
	var brand *api_models.NameReference
	if model.Brand != nil {
		brandID := model.Brand.ID.String()
		brandName := model.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	return &api_models.Model{
		Id:    model.ID.String(),
		Name:  model.Name,
		Brand: brand,
	}, nil
}

func UpdateModelService(
	ctx context.Context,
	id string,
	request api_models.UpdateModelRequest,
) (*api_models.Model, error) {

	modelID, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	validationResult := validation.ValidateUpdateModelRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	var (
		brandID           *uuid.UUID
		shouldUpdateBrand bool
	)

	if request.Brand != nil {

		shouldUpdateBrand = true

		if !IsEmptyNameReference(request.Brand) {

			brandID, err = ResolveBrandID(
				ctx,
				request.Brand,
			)

			if err != nil {
				return nil, err
			}
		}
	}

	repo := repositories.NewModelRepository()

	model, err := repo.UpdateRepository(
		ctx,
		modelID,
		request,
		brandID,
		shouldUpdateBrand,
	)

	if err != nil {
		return nil, err
	}

	model, err = repo.GetByIDRepository(
		ctx,
		model.ID,
	)

	if err != nil {
		return nil, err
	}

	var brand *api_models.NameReference
	if model.Brand != nil {
		brandID := model.Brand.ID.String()
		brandName := model.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	return &api_models.Model{
		Id:    model.ID.String(),
		Name:  model.Name,
		Brand: brand,
	}, nil
}
