//go:build integration || e2e

package tests

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	if err := PrepareDatabase(); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	if err := CleanupDatabase(); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}
