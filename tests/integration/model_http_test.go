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
	model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
)

func Test_Model_HTTP_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/models",
		nil,
	)

	rec := httptest.NewRecorder()

	model.ModelsMainHandler(
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

func Test_Model_HTTP_Create(t *testing.T) {
	brandCreated, err := brand.CreateBrandService(
		context.Background(),
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"TOYOTA_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	request := api_models.CreateModelRequest{
		Name: fmt.Sprintf(
			"CAMRY_%d",
			time.Now().UnixNano(),
		),
		Brand: &api_models.NameReference{
			Id: &brandCreated.Id,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/models",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	model.ModelsMainHandler(
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

	var response api_models.Model

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

	if response.Brand == nil {
		t.Fatal("brand отсутствует")
	}
}

func Test_Model_HTTP_Create_WithNewBrand(t *testing.T) {
	brandName := fmt.Sprintf(
		"LADA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"GRANTA_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateModelRequest{
		Name: modelName,
		Brand: &api_models.NameReference{
			Name: &brandName,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/models",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	model.ModelsMainHandler(
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

	var response api_models.Model

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Brand == nil {
		t.Fatal("brand отсутствует")
	}

	if response.Brand.Name == nil {
		t.Fatal("brand.name отсутствует")
	}

	if *response.Brand.Name != brandName {
		t.Fatalf(
			"ожидался бренд %q, получен %q",
			brandName,
			*response.Brand.Name,
		)
	}
}

func Test_Model_HTTP_Update(t *testing.T) {
	ctx := context.Background()

	brandCreated, err := brand.CreateBrandService(
		ctx,
		api_models.CreateNameEntityRequest{
			Name: fmt.Sprintf(
				"BMW_%d",
				time.Now().UnixNano(),
			),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	created, err := model.CreateModelService(
		ctx,
		api_models.CreateModelRequest{
			Name: fmt.Sprintf(
				"MODEL_%d",
				time.Now().UnixNano(),
			),
			Brand: &api_models.NameReference{
				Id: &brandCreated.Id,
			},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	updatedName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateModelRequest{
		Name: &updatedName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/models/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	model.ModelByIDMainHandler(
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

	var response api_models.Model

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

func Test_Model_HTTP_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/models/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	model.ModelByIDMainHandler(
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

func Test_Model_HTTP_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/models",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	model.ModelsMainHandler(
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

func Test_Model_HTTP_Create_EmptyName(t *testing.T) {
	request := api_models.CreateModelRequest{}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/models",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	model.ModelsMainHandler(
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

func Test_Model_HTTP_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/models/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	model.ModelByIDMainHandler(
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
