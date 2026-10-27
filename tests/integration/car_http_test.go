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
	car "github.com/Adisey/car-rental-api/internal/cars"
)

func Test_Car_HTTP_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/cars",
		nil,
	)

	rec := httptest.NewRecorder()

	car.CarsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}

func Test_Car_HTTP_Create(t *testing.T) {
	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateCarRequest{
		Name: "BMW X5",

		Brand: &api_models.NameReference{
			Name: &brandName,
		},

		Model: &api_models.CarModel{
			Name: &modelName,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/cars",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	car.CarsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Car

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Id == "" {
		t.Fatal("empty id")
	}

	if response.Brand == nil {
		t.Fatal("brand missing")
	}

	if response.Model == nil {
		t.Fatal("model missing")
	}
}

func Test_Car_HTTP_Create_WithNewReferences(t *testing.T) {
	brandName := fmt.Sprintf(
		"LADA_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"GRANTA_%d",
		time.Now().UnixNano(),
	)

	fuelCode := fmt.Sprintf(
		"PETROL_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateCarRequest{
		Name: "Lada Granta",

		Brand: &api_models.NameReference{
			Name: &brandName,
		},

		Model: &api_models.CarModel{
			Name: &modelName,
		},

		FuelType: &api_models.CodeReference{
			Code: &fuelCode,
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/cars",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	car.CarsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusCreated,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Car

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Brand == nil ||
		response.Brand.Name == nil ||
		*response.Brand.Name != brandName {

		t.Fatal("invalid brand")
	}

	if response.Model == nil ||
		response.Model.Name != modelName {

		t.Fatal("invalid model")
	}

	if response.FuelType == nil ||
		response.FuelType.Code == nil ||
		*response.FuelType.Code != fuelCode {

		t.Fatal("invalid fuel type")
	}
}

func Test_Car_HTTP_Update(t *testing.T) {
	ctx := context.Background()

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Original Car",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"UPDATED_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateCarRequest{
		Name: &newName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/cars/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	car.CarByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Car

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Name != newName {
		t.Fatal("name not updated")
	}
}

func Test_Car_HTTP_Update_NameOnly_KeepsReferences(t *testing.T) {
	ctx := context.Background()

	brandName := fmt.Sprintf(
		"BMW_%d",
		time.Now().UnixNano(),
	)

	modelName := fmt.Sprintf(
		"X5_%d",
		time.Now().UnixNano(),
	)

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "BMW X5",

			Brand: &api_models.NameReference{
				Name: &brandName,
			},

			Model: &api_models.CarModel{
				Name: &modelName,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"RESTYLE_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateCarRequest{
		Name: &newName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/cars/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	car.CarByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusOK,
			rec.Code,
			rec.Body.String(),
		)
	}

	var response api_models.Car

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Brand == nil {
		t.Fatal("brand disappeared")
	}

	if response.Model == nil {
		t.Fatal("model disappeared")
	}

	actual, err := car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.Brand == nil {
		t.Fatal("brand disappeared in db")
	}

	if actual.Model == nil {
		t.Fatal("model disappeared in db")
	}
}

func Test_Car_HTTP_Delete(t *testing.T) {
	ctx := context.Background()

	created, err := car.CreateCarService(
		ctx,
		api_models.CreateCarRequest{
			Name: "Delete Via HTTP",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodDelete,
		"/cars/"+created.Id,
		nil,
	)

	rec := httptest.NewRecorder()

	car.CarByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusNoContent &&
		rec.Code != http.StatusOK {

		t.Fatalf(
			"unexpected status %d body=%s",
			rec.Code,
			rec.Body.String(),
		)
	}

	_, err = car.GetCarByIDService(
		ctx,
		created.Id,
	)

	if err == nil {
		t.Fatal("car still exists")
	}
}

func Test_Car_HTTP_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/cars/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	car.CarByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Car_HTTP_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/cars",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	car.CarsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Car_HTTP_Create_EmptyName(t *testing.T) {
	request := api_models.CreateCarRequest{}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/cars",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	car.CarsMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}

func Test_Car_HTTP_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/cars/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	car.CarByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			rec.Code,
		)
	}
}
