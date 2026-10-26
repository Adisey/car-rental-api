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
	brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
)

func Test_HTTP_Brand_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/brands",
		nil,
	)

	rec := httptest.NewRecorder()

	brand.BrandsMainHandler(
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

func Test_HTTP_Brand_Create(t *testing.T) {
	request := api_models.CreateBrandRequest{
		Name: fmt.Sprintf(
			"HTTP_BRAND_%d",
			time.Now().UnixNano(),
		),
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/brands",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	brand.BrandsMainHandler(
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

	var response api_models.Brand

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name != request.Name {
		t.Fatalf(
			"ожидалось имя %q, получено %q",
			request.Name,
			response.Name,
		)
	}
}

func Test_HTTP_Brand_Update(t *testing.T) {
	ctx := context.Background()

	created, err := brand.CreateBrandService(
		ctx,
		api_models.CreateBrandRequest{
			Name: fmt.Sprintf(
				"PATCH_BRAND_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	updatedName := fmt.Sprintf(
		"PATCH_UPDATED_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateBrandRequest{
		Name: &updatedName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/brands/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	brand.BrandByIDMainHandler(
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

	var response api_models.Brand

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name != updatedName {
		t.Fatalf(
			"ожидалось имя %q, получено %q",
			updatedName,
			response.Name,
		)
	}
}

func Test_HTTP_Brand_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/brands/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	brand.BrandByIDMainHandler(
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

func Test_HTTP_Brand_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/brands",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	brand.BrandsMainHandler(
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

func Test_HTTP_Brand_Create_EmptyName(t *testing.T) {
	request := api_models.CreateBrandRequest{}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/brands",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	brand.BrandsMainHandler(
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

func Test_HTTP_Brand_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/brands/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	brand.BrandByIDMainHandler(
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
