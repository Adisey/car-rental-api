//go:build e2e

package e2e

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func Test_GetCars(t *testing.T) {

	baseURL := os.Getenv("TEST_BASE_URL")

	resp, err := http.Get(
		baseURL + "/cars",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			resp.StatusCode,
		)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(
		"cars response: %s",
		string(body),
	)
}