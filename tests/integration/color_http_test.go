//go:build integration

package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
)

func Test_Color_HTTP_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/colors",
		nil,
	)

	rec := httptest.NewRecorder()

	color.ColorsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"ожидался статус %d, получен %d",
			http.StatusOK,
			rec.Code,
		)
	}
}

func Test_Color_HTTP_Create(t *testing.T) {
	hex := "#FF0000"

	request := api_models.CreateColorRequest{
		Code: fmt.Sprintf(
			"HTTP_CREATE_%d",
			time.Now().UnixNano(),
		),
		Hex: &hex,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/colors",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	color.ColorsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"ожидался статус %d, получен %d\nbody=%s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Color

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Code != request.Code {
		t.Fatalf(
			"ожидался код %q, получен %q",
			request.Code,
			response.Code,
		)
	}

	if response.Hex == nil {
		t.Fatal("hex отсутствует в ответе")
	}

	if *response.Hex != hex {
		t.Fatalf(
			"ожидался hex %q, получен %q",
			hex,
			*response.Hex,
		)
	}
}

func Test_Color_HTTP_Update(t *testing.T) {
	ctx := context.Background()

	created, err := color.CreateColorService(
		ctx,
		api_models.CreateColorRequest{
			Code: fmt.Sprintf(
				"PATCH_TEST_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	updatedCode := fmt.Sprintf(
		"PATCH_UPDATED_%d",
		time.Now().UnixNano(),
	)

	updatedHex := "#00FF00"

	request := api_models.UpdateColorRequest{
		Code: &updatedCode,
		Hex:  &updatedHex,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/colors/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	color.ColorByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"ожидался статус %d, получен %d\nbody=%s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Color

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Code != updatedCode {
		t.Fatalf(
			"ожидался код %q, получен %q",
			updatedCode,
			response.Code,
		)
	}

	if response.Hex == nil {
		t.Fatal("hex отсутствует после обновления")
	}

	if *response.Hex != updatedHex {
		t.Fatalf(
			"ожидался hex %q, получен %q",
			updatedHex,
			*response.Hex,
		)
	}
}

func Test_Color_HTTP_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/colors/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	color.ColorByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"ожидался статус %d, получен %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Color_HTTP_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/colors",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	color.ColorsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"ожидался статус %d, получен %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Color_HTTP_Create_EmptyCode(t *testing.T) {
	request := api_models.CreateColorRequest{
		Code: "",
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/colors",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	color.ColorsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"ожидался статус %d, получен %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Color_HTTP_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/colors/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	color.ColorByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"ожидался статус %d, получен %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
