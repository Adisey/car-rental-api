//go:build integration

package integration

import (
	"fmt"
	"log"
	"os"
	"testing"

	db "github.com/Adisey/car-rental-api/internal/db"
	tests "github.com/Adisey/car-rental-api/tests"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	println("TestMainIntegration - Load env.test")
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Fatal(err)
	}
	println("TestMainIntegration - Will RUN PrepareDatabase")
	if err := tests.PrepareDatabase(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("TestMainIntegration - Will Connect DB")
	if err := db.ConnectBun(); err != nil {
		log.Fatal(err)
	}
	println("TestMainIntegration - Will RUN Tests")
	code := m.Run()

	println("TestMainIntegration - Will Exit")
	os.Exit(code)
}
