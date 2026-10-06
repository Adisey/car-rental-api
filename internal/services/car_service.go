package services

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/db_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
	"github.com/Adisey/car-rental-api/internal/validation"
)

var cars = []api_models.Car{
	{
		Id:   "1",
		Name: "BMW X5",
	},
	{
		Id:   "2",
		Name: "Toyota Corolla",
	},
	{
		Id:   "3",
		Name: "Skoda Octavia",
	},
}

func GetCarsService(ctx context.Context) ([]api_models.Car, error) {
	repo := repositories.NewCarRepository()

	dbCars, err := repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	cars := make([]api_models.Car, 0, len(dbCars))

	for _, car := range dbCars {
		cars = append(cars, api_models.Car{
			Id:          car.ID.String(),
			Name:        car.Name,
			Description: car.Description,
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
	log.Println("Handler GET /cars/", id)

	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbCar, err := repo.GetByIDRepository(ctx, carID)
	if err != nil {
		return nil, err
	}

	return &api_models.Car{
		Id:          dbCar.ID.String(),
		Name:        dbCar.Name,
		Description: dbCar.Description,
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

	repo := repositories.NewCarRepository()

	car := &db_models.Car{
		Name:        request.Name,
		Description: request.Description,
	}

	err := repo.CreateRepository(ctx, car)
	if err != nil {
		return nil, err
	}

	return &api_models.Car{
		Id:          car.ID.String(),
		Name:        car.Name,
		Description: car.Description,
		CreatedAt:   car.CreatedAt,
		UpdatedAt:   car.UpdatedAt,
	}, nil
}

func UpdateCar(
	id string,
	request api_models.UpdateCarRequest,
) (*api_models.Car, bool) {

	for i, car := range cars {
		if car.Id == id {
			cars[i].Name = request.Name
			cars[i].Description = request.Description

			return &cars[i], true
		}
	}

	return nil, false
}

func DeleteCar(id string) bool {
	for i, car := range cars {
		if car.Id == id {
			cars = append(cars[:i], cars[i+1:]...)
			return true
		}
	}

	return false
}
