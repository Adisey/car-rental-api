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
	body_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
)

/*
Проверяет получение списка типов кузова через HTTP.

Цепочка:

HTTP -> Handler -> Service -> Repository -> PostgreSQL
*/
func Test_BodyType_HTTP_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/bodyTypes",
		nil,
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypesMainHandler(
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

	if rec.Body.Len() == 0 {
		t.Fatal("ожидалось непустое тело ответа")
	}
}

/*
Проверяет создание записи через HTTP POST.
*/
func Test_BodyType_HTTP_Create(t *testing.T) {
	request := api_models.CreateBodyTypeRequest{
		Code: fmt.Sprintf(
			"HTTP_CREATE_%d",
			time.Now().UnixNano(),
		),
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/bodyTypes",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypesMainHandler(
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

	var response api_models.BodyType

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Id == "" {
		t.Fatal("не получен идентификатор созданной записи")
	}

	if response.Code != request.Code {
		t.Fatalf(
			"ожидался код %q, получен %q",
			request.Code,
			response.Code,
		)
	}
}

/*
Проверяет обновление записи через HTTP PATCH.
*/
func Test_BodyType_HTTP_Update(t *testing.T) {
	ctx := context.Background()

	created, err := body_type.CreateBodyTypeService(
		ctx,
		api_models.CreateCodeEntityRequest{
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

	request := api_models.UpdateBodyTypeRequest{
		Code: &updatedCode,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/bodyTypes/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypeByIDMainHandler(
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

	var response api_models.BodyType

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
}

func Test_BodyType_HTTP_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/bodyTypes/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypeByIDMainHandler(
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

func Test_BodyType_HTTP_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/bodyTypes",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypesMainHandler(
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
func Test_BodyType_HTTP_Create_EmptyCode(t *testing.T) {
	request := api_models.CreateBodyTypeRequest{
		Code: "",
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/bodyTypes",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypesMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"ожидался статус %d, получен %d\nbody=%s",
			http.StatusBadRequest,
			rec.Code,
			rec.Body.String(),
		)
	}
}
func Test_BodyType_HTTP_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/bodyTypes/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	body_type.BodyTypeByIDMainHandler(
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
