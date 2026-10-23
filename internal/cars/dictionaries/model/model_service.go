package car_model

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	shared_references "github.com/Adisey/car-rental-api/internal/shared/references"
)

func GetModelsService(
	ctx context.Context,
	brandID *uuid.UUID,
) ([]api_models.Model, error) {

	repo := NewModelRepository()

	dbModels, err := repo.GetModelAllRepository(
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
	println("----modelID->", modelID.String())

	if err != nil {
		return nil, err
	}

	repoModel := NewModelRepository()
	println("----repoModel->", repoModel)

	dbModel, err := repoModel.GetModelByIDRepository(
		ctx,
		modelID,
	)
	println("----dbModel->", dbModel)
	println("----dbModel-err->", err)
	if err != nil {
		return nil, err
	}

	var brand *api_models.NameReference

	if dbModel.Brand != nil {
		brandID := dbModel.Brand.ID.String()
		println("----dbModel-brandID->", brandID)
		brandName := dbModel.Brand.Name
		println("----dbModel-brandName->", brandName)
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	println("----dbModel-brand->", brand)

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

	validationResult := ValidateCreateModelRequest(
		request,
	)

	if validationResult.HasErrors() {
		return nil, validationResult
	}

	var brandID *uuid.UUID

	if request.Brand != nil &&
		!shared_references.IsEmptyNameReference(request.Brand) {

		var err error

		brandID, err = car_brand.ResolveBrandID(

			ctx,
			request.Brand,
		)

		if err != nil {
			return nil, err
		}
	}

	repo := NewModelRepository()

	existingModel, err := repo.GetModelByNameRepository(
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

	model := &ModelDBModel{
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

	model, err = repo.GetModelByIDRepository(
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

	validationResult := ValidateUpdateModelRequest(
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

		if !shared_references.IsEmptyNameReference(request.Brand) {

			brandID, err = car_brand.ResolveBrandID(
				ctx,
				request.Brand,
			)

			if err != nil {
				return nil, err
			}
		}
	}

	repo := NewModelRepository()

	model, err := repo.UpdateModelRepository(
		ctx,
		modelID,
		request,
		brandID,
		shouldUpdateBrand,
	)

	if err != nil {
		return nil, err
	}

	model, err = repo.GetModelByIDRepository(
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
