package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

func GetCarsService(ctx context.Context) ([]api_models.Car, error) {
	repo := repositories.NewCarRepository()
	dbCars, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	cars := make([]api_models.Car, 0, len(dbCars))
	for _, car := range dbCars {
		var color *api_models.CodeReference
		if car.Color != nil {
			colorID := car.Color.ID.String()
			colorCode := car.Color.Code

			color = &api_models.CodeReference{
				Id:   &colorID,
				Code: &colorCode,
			}
		}
		var brand *api_models.NameReference
		if car.Brand != nil {
			brandID := car.Brand.ID.String()
			brandCode := car.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandCode,
			}
		}
		var model *api_models.Model
		if car.Model != nil {
			modelID := car.Model.ID.String()
			modelCode := car.Model.Name
			brandID := car.Model.Brand.ID.String()
			brandCode := car.Model.Brand.Name
			brand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandCode,
			}
			model = &api_models.Model{
				Id:    modelID,
				Name:  modelCode,
				Brand: brand,
			}
		}

		cars = append(cars, api_models.Car{
			Id:          car.ID.String(),
			Name:        car.Name,
			Description: car.Description,
			Color:       color,
			Brand:       brand,
			Model:       model,
			CreatedAt:   car.CreatedAt,
			UpdatedAt:   car.UpdatedAt,
		})
	}
	return cars, nil
}

func GetCarByIDService(
	ctx context.Context,
	id string,
) (*api_models.Car, error) {
	repo := repositories.NewCarRepository()
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	dbCar, err := repo.GetByIDRepository(ctx, carID)
	if err != nil {
		return nil, err
	}
	var color *api_models.CodeReference
	if dbCar.Color != nil {
		colorID := dbCar.Color.ID.String()
		colorCode := dbCar.Color.Code
		color = &api_models.CodeReference{
			Id:   &colorID,
			Code: &colorCode,
		}
	}
	var brand *api_models.NameReference
	if dbCar.Brand != nil {
		brandID := dbCar.Brand.ID.String()
		brandCode := dbCar.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandCode,
		}
	}
	var model *api_models.Model
	if dbCar.Model != nil {
		modelID := dbCar.Model.ID.String()
		modelCode := dbCar.Model.Name
		brandID := dbCar.Model.Brand.ID.String()
		brandCode := dbCar.Model.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandCode,
		}
		model = &api_models.Model{
			Id:    modelID,
			Name:  modelCode,
			Brand: brand,
		}
	}

	return &api_models.Car{
		Id:          dbCar.ID.String(),
		Name:        dbCar.Name,
		Description: dbCar.Description,
		Color:       color,
		Brand:       brand,
		Model:       model,
		CreatedAt:   dbCar.CreatedAt,
		UpdatedAt:   dbCar.UpdatedAt,
	}, nil
}

func CreateCarService(
	ctx context.Context,
	request api_models.CreateCarRequest,
) (*api_models.Car, error) {
	validationResult := validation.ValidateCreateCarRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	var colorID *uuid.UUID
	if request.Color != nil &&
		!IsEmptyColor(request.Color) {
		var err error
		colorID, err = ResolveColorID(
			ctx,
			request.Color,
		)
		if err != nil {
			return nil, err
		}
	}
	var brandID *uuid.UUID
	if request.Brand != nil &&
		!IsEmptyBrand(request.Brand) {
		var err error
		brandID, err = ResolveBrandID(
			ctx,
			request.Brand,
		)
		if err != nil {
			return nil, err
		}
	}
	var modelID *uuid.UUID
	if request.Model != nil {
		model := request.Model
		if model.Brand == nil &&
			brandID != nil {
			brandIDString := brandID.String()
			model.Brand = &api_models.NameReference{
				Id: &brandIDString,
			}
		}
		if !IsEmptyModel(model) {
			var err error
			modelID, err = ResolveModelID(
				ctx,
				model,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	repo := repositories.NewCarRepository()
	car := &db_models.Car{
		Name:        request.Name,
		Description: request.Description,
		ColorID:     colorID,
		BrandID:     brandID,
		ModelID:     modelID,
	}
	err := repo.CreateRepository(ctx, car)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)

	if err != nil {
		return nil, err
	}
	var color *api_models.CodeReference
	if car.Color != nil {
		colorID := car.Color.ID.String()
		colorCode := car.Color.Code
		color = &api_models.CodeReference{
			Id:   &colorID,
			Code: &colorCode,
		}
	}
	var brand *api_models.NameReference
	if car.Brand != nil {
		brandID := car.Brand.ID.String()
		brandName := car.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	var model *api_models.Model
	if car.Model != nil {
		var modelBrand *api_models.NameReference
		if car.Model.Brand != nil {
			brandID := car.Model.Brand.ID.String()
			brandName := car.Model.Brand.Name
			modelBrand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}
		newModelID := car.Model.ID.String()
		newModelName := car.Model.Name
		model = &api_models.Model{
			Id:    newModelID,
			Name:  newModelName,
			Brand: modelBrand,
		}
	}
	return &api_models.Car{
		Id:          car.ID.String(),
		Name:        car.Name,
		Description: car.Description,
		Color:       color,
		Brand:       brand,
		Model:       model,
		CreatedAt:   car.CreatedAt,
		UpdatedAt:   car.UpdatedAt,
	}, nil
}

func UpdateCarService(
	ctx context.Context,
	id string,
	request api_models.UpdateCarRequest,
) (*api_models.Car, error) {
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	validationResult := validation.ValidateUpdateCarRequest(
		request,
	)
	if validationResult.HasErrors() {
		return nil, validationResult
	}
	var (
		colorID           *uuid.UUID
		shouldUpdateColor bool
	)
	if request.Color != nil {
		shouldUpdateColor = true
		if !IsEmptyColor(request.Color) {
			colorID, err = ResolveColorID(
				ctx,
				request.Color,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		brandID           *uuid.UUID
		shouldUpdateBrand bool
	)
	if request.Brand != nil {
		shouldUpdateBrand = true
		if !IsEmptyBrand(request.Brand) {
			brandID, err = ResolveBrandID(
				ctx,
				request.Brand,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	var (
		modelID           *uuid.UUID
		shouldUpdateModel bool
	)
	if request.Model != nil {
		shouldUpdateModel = true
		model := request.Model
		if model.Brand == nil &&
			brandID != nil {
			brandIDString := brandID.String()
			model.Brand = &api_models.NameReference{
				Id: &brandIDString,
			}
		}
		if !IsEmptyModel(model) {
			var err error
			modelID, err = ResolveModelID(
				ctx,
				model,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	repo := repositories.NewCarRepository()
	car, err := repo.UpdateRepository(
		ctx,
		carID,
		request,
		colorID,
		shouldUpdateColor,
		brandID,
		shouldUpdateBrand,
		modelID,
		shouldUpdateModel,
	)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)
	if err != nil {
		return nil, err
	}
	var color *api_models.CodeReference
	if car.Color != nil {
		colorID := car.Color.ID.String()
		colorCode := car.Color.Code
		color = &api_models.CodeReference{
			Id:   &colorID,
			Code: &colorCode,
		}
	}
	var brand *api_models.NameReference
	if car.Brand != nil {
		brandID := car.Brand.ID.String()
		brandName := car.Brand.Name
		brand = &api_models.NameReference{
			Id:   &brandID,
			Name: &brandName,
		}
	}
	var model *api_models.Model
	if car.Model != nil {
		var modelBrand *api_models.NameReference
		if car.Model.Brand != nil {
			brandID := car.Model.Brand.ID.String()
			brandName := car.Model.Brand.Name
			modelBrand = &api_models.NameReference{
				Id:   &brandID,
				Name: &brandName,
			}
		}
		newModelID := car.Model.ID.String()
		newModelName := car.Model.Name
		model = &api_models.Model{
			Id:    newModelID,
			Name:  newModelName,
			Brand: modelBrand,
		}
	}
	return &api_models.Car{
		Id:          car.ID.String(),
		Name:        car.Name,
		Description: car.Description,
		Color:       color,
		Brand:       brand,
		Model:       model,
		CreatedAt:   car.CreatedAt,
		UpdatedAt:   car.UpdatedAt,
	}, nil
}

func DeleteCarService(ctx context.Context, id string) error {

	carID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	repo := repositories.NewCarRepository()

	return repo.DeleteRepository(
		ctx,
		carID,
	)
}
