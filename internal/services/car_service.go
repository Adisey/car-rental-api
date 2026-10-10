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

		var color *api_models.Color
		if car.Color != nil {
			color = &api_models.Color{
				Id:   car.Color.ID.String(),
				Code: car.Color.Code,
			}
		}

		cars = append(cars, api_models.Car{
			Id:          car.ID.String(),
			Name:        car.Name,
			Description: car.Description,
			Color:       color,
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

	var color *api_models.Color

	if dbCar.Color != nil {
		color = &api_models.Color{
			Id:   dbCar.Color.ID.String(),
			Code: dbCar.Color.Code,
		}
	}

	return &api_models.Car{
		Id:          dbCar.ID.String(),
		Name:        dbCar.Name,
		Description: dbCar.Description,
		Color:       color,
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
	repo := repositories.NewCarRepository()
	car := &db_models.Car{
		Name:        request.Name,
		Description: request.Description,
		ColorID:     colorID,
		BrandID:     brandID,
	}
	err := repo.CreateRepository(ctx, car)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)

	if err != nil {
		return nil, err
	}
	var color *api_models.Color
	if car.Color != nil {
		color = &api_models.Color{
			Id:   car.Color.ID.String(),
			Code: car.Color.Code,
		}
	}
	var brand *api_models.Brand
	if car.Brand != nil {
		brand = &api_models.Brand{
			Id:   car.Brand.ID.String(),
			Name: car.Brand.Name,
		}
	}
	return &api_models.Car{
		Id:          car.ID.String(),
		Name:        car.Name,
		Description: car.Description,
		Color:       color,
		Brand:       brand,
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
	repo := repositories.NewCarRepository()
	car, err := repo.UpdateRepository(
		ctx,
		carID,
		request,
		colorID,
		shouldUpdateColor,
		brandID,
		shouldUpdateBrand,
	)
	if err != nil {
		return nil, err
	}
	car, err = repo.GetByIDRepository(ctx, car.ID)
	if err != nil {
		return nil, err
	}
	var color *api_models.Color
	if car.Color != nil {
		color = &api_models.Color{
			Id:   car.Color.ID.String(),
			Code: car.Color.Code,
		}
	}
	var brand *api_models.Brand
	if car.Brand != nil {
		brand = &api_models.Brand{
			Id:   car.BrandID.String(),
			Name: car.Brand.Name,
		}
	}
	return &api_models.Car{
		Id:          car.ID.String(),
		Name:        car.Name,
		Description: car.Description,
		Color:       color,
		Brand:       brand,
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
