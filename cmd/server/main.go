package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	car "github.com/Adisey/car-rental-api/internal/cars"
	car_body_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
	car_model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
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
	http.HandleFunc("/cars", car.CarsMainHandler)
	http.HandleFunc("/cars/", car.CarByIDMainHandler)
	http.HandleFunc("/colors", car_color.ColorsMainHandler)
	http.HandleFunc("/colors/", car_color.ColorByIDMainHandler)
	http.HandleFunc("/body_types", car_body_type.BodyTypesMainHandler)
	http.HandleFunc("/body_types/", car_body_type.BodyTypeByIDMainHandler)
	http.HandleFunc("/brands", car_brand.BrandsMainHandler)
	http.HandleFunc("/brands/", car_brand.BrandByIDMainHandler)
	http.HandleFunc("/models", car_model.ModelsMainHandler)
	http.HandleFunc("/models/", car_model.ModelByIDMainHandler)

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
