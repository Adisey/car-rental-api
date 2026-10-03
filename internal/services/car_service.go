package services

import api_models "github.com/Adisey/car-rental-api/internal/api_models"

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

func GetCars() []api_models.Car {
	return cars
}

func GetCarByID(id string) (*api_models.Car, bool) {
	for _, car := range cars {
		if car.Id == id {
			return &car, true
		}
	}

	return nil, false
}
