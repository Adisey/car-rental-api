package main

import (
	"fmt"
	"os"

	"ariga.io/atlas-provider-bun/bunschema"

	"github.com/Adisey/car-rental-api/internal/db_models"
)

func main() {
	stmts, err := bunschema.New(
		bunschema.DialectPostgres,
	).Load(
		&db_models.Car{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(stmts)
}
