package services

import (
	"context"
	"strconv"

	"github.com/Adisey/car-rental-api/internal/api_models"
	"github.com/Adisey/car-rental-api/internal/repositories"
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

func GetCars(ctx context.Context) ([]api_models.Car, error) {
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
		})
	}

	return cars, nil
}

func GetCarByID(id string) (*api_models.Car, bool) {
	for _, car := range cars {
		if car.Id == id {
			return &car, true
		}
	}

	return nil, false
}

func CreateCar(request api_models.CreateCarRequest) api_models.Car {
	car := api_models.Car{
		Id:          strconv.Itoa(len(cars) + 1),
		Name:        request.Name,
		Description: request.Description,
	}

	cars = append(cars, car)

	return car
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
