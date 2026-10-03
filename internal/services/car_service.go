package services

import "github.com/Adisey/car-rental-api/internal/models"

var cars = []models.Car{
	{
		ID:   "1",
		Name: "BMW X5",
	},
	{
		ID:   "2",
		Name: "Toyota Corolla",
	},
	{
		ID:   "3",
		Name: "Skoda Octavia",
	},
}

func GetCars() []models.Car {
	return cars
}

func GetCarByID(id string) (*models.Car, bool) {
	for _, car := range cars {
		if car.ID == id {
			return &car, true
		}
	}

	return nil, false
}
