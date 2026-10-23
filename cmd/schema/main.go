package main

import (
	"fmt"
	"os"

	"ariga.io/atlas-provider-bun/bunschema"
	car "github.com/Adisey/car-rental-api/internal/cars"
	car_body_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
)

func main() {
	stmts, err := bunschema.New(
		bunschema.DialectPostgres,
	).Load(
		&car.Car{},
		&car_color.Color{},
		&car_body_type.BodyType{},
		&car_brand.Brand{},
		&db_models.DriveType{},
		&db_models.FuelType{},
		&db_models.Model{},
		&db_models.TransmissionType{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(stmts)
}
