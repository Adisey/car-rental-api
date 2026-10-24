package main

import (
	"fmt"
	"os"

	"ariga.io/atlas-provider-bun/bunschema"
	car "github.com/Adisey/car-rental-api/internal/cars"
	car_body "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
	car_fuel "github.com/Adisey/car-rental-api/internal/cars/dictionaries/fuel_type"
	car_model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
	car_transmission "github.com/Adisey/car-rental-api/internal/cars/dictionaries/transmission_type"
)

func main() {
	stmts, err := bunschema.New(
		bunschema.DialectPostgres,
	).Load(
		&car.CarDBModel{},
		&car_color.ColorDBModel{},
		&car_body.BodyTypeDBModel{},
		&car_brand.BrandDBModel{},
		&car_body.BodyTypeDBModel{},
		&car_fuel.FuelTypeDBModel{},
		&car_model.ModelDBModel{},
		&car_transmission.TransmissionTypeDBModel{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(stmts)
}
