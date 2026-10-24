//go:build integration || e2e

package tests

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	println("TestMain - Load env.test")
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Fatal(err)
	}
	println("TestMain - Will RUN PrepareDatabase")
	if err := PrepareDatabase(); err != nil {
		log.Fatal(err)
	}
	println("TestMain - Will RUN Tests")
	code := m.Run()

	println("Will Exit")
	os.Exit(code)
}
