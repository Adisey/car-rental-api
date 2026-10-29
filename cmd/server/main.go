package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	cars "github.com/Adisey/car-rental-api/internal/cars"
	car_body "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
	car_drive "github.com/Adisey/car-rental-api/internal/cars/dictionaries/drive_type"
	car_fuel "github.com/Adisey/car-rental-api/internal/cars/dictionaries/fuel_type"
	car_model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
	car_transmission "github.com/Adisey/car-rental-api/internal/cars/dictionaries/transmission_type"
	customers "github.com/Adisey/car-rental-api/internal/customers"

	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/health"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println(".env.local file not found")
	}
	ctx := context.Background()
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	if err := db.ConnectBun(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", health.Health)

	http.HandleFunc("/cars", cars.CarsMainHandler)
	http.HandleFunc("/cars/", cars.CarByIDMainHandler)

	http.HandleFunc("/body_types", car_body.BodyTypesMainHandler)
	http.HandleFunc("/body_types/", car_body.BodyTypeByIDMainHandler)

	http.HandleFunc("/brands", car_brand.BrandsMainHandler)
	http.HandleFunc("/brands/", car_brand.BrandByIDMainHandler)

	http.HandleFunc("/colors", car_color.ColorsMainHandler)
	http.HandleFunc("/colors/", car_color.ColorByIDMainHandler)

	http.HandleFunc("/drive_type", car_drive.DriveTypesMainHandler)
	http.HandleFunc("/drive_type/", car_drive.DriveTypeByIDMainHandler)

	http.HandleFunc("/fuel_type", car_fuel.FuelTypesMainHandler)
	http.HandleFunc("/fuel_type/", car_fuel.FuelTypeByIDMainHandler)

	http.HandleFunc("/models", car_model.ModelsMainHandler)
	http.HandleFunc("/models/", car_model.ModelByIDMainHandler)

	http.HandleFunc("/transmission_type", car_transmission.TransmissionTypesMainHandler)
	http.HandleFunc("/transmission_type/", car_transmission.TransmissionTypeByIDMainHandler)

	http.HandleFunc("/customers", customers.CustomersMainHandler)
	http.HandleFunc("/customers/", customers.CustomerByIDMainHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
