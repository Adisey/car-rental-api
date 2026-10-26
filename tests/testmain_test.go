package tests

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("TestMain - Will RUN Tests")
	code := m.Run()

	println("TestMain - Will Exit")
	os.Exit(code)
}
