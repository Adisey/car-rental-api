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
	customers "github.com/Adisey/car-rental-api/internal/customers"
)

func Test_Customer_HTTP_GetAll(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/customers",
		nil,
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
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

func Test_Customer_HTTP_Create(t *testing.T) {
	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	email := fmt.Sprintf(
		"ivan_%d@test.com",
		time.Now().UnixNano(),
	)

	request := api_models.CreateCustomerRequest{
		FirstName: firstName,
		Phone:     &phone,
		Email:     &email,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
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

	var response api_models.Customer

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Id == "" {
		t.Fatal("empty id")
	}

	if response.FirstName != firstName {
		t.Fatal("invalid first_name")
	}

	if response.Phone == nil {
		t.Fatal("phone missing")
	}

	if *response.Phone != phone {
		t.Fatal("invalid phone")
	}

	if response.Email == nil {
		t.Fatal("email missing")
	}

	if *response.Email != email {
		t.Fatal("invalid email")
	}
}

func Test_Customer_HTTP_Update(t *testing.T) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateCustomerRequest{
		FirstName: &newName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/customers/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
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

	var response api_models.Customer

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.FirstName != newName {
		t.Fatal("first_name not updated")
	}
}

func Test_Customer_HTTP_Update_FirstNameOnly_KeepsOtherFields(
	t *testing.T,
) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	lastName := fmt.Sprintf(
		"Ivanov_%d",
		time.Now().UnixNano(),
	)

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	email := fmt.Sprintf(
		"ivan_%d@test.com",
		time.Now().UnixNano(),
	)

	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			LastName:  &lastName,
			Phone:     &phone,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateCustomerRequest{
		FirstName: &newName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/customers/"+created.Id,
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
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

	actual, err := customers.GetCustomerByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.FirstName != newName {
		t.Fatal("first_name not updated")
	}

	if actual.LastName == nil {
		t.Fatal("last_name disappeared")
	}

	if *actual.LastName != lastName {
		t.Fatal("last_name changed")
	}

	if actual.Phone == nil {
		t.Fatal("phone disappeared")
	}

	if *actual.Phone != phone {
		t.Fatal("phone changed")
	}

	if actual.Email == nil {
		t.Fatal("email disappeared")
	}

	if *actual.Email != email {
		t.Fatal("email changed")
	}
}

func Test_Customer_HTTP_Create_DuplicatePhone(t *testing.T) {
	ctx := context.Background()

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	_, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			Phone:     &phone,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	request := api_models.CreateCustomerRequest{
		FirstName: fmt.Sprintf(
			"Petr_%d",
			time.Now().UnixNano(),
		),
		Phone: &phone,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
		rec,
		req,
	)

	if rec.Code == http.StatusCreated {
		t.Fatal("expected duplicate phone error")
	}
}

func Test_Customer_HTTP_Create_DuplicateEmail(t *testing.T) {
	ctx := context.Background()

	email := fmt.Sprintf(
		"mail_%d@test.com",
		time.Now().UnixNano(),
	)

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	_, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	request := api_models.CreateCustomerRequest{
		FirstName: fmt.Sprintf(
			"Petr_%d",
			time.Now().UnixNano(),
		),
		Email: &email,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
		rec,
		req,
	)

	if rec.Code == http.StatusCreated {
		t.Fatal("expected duplicate email error")
	}
}

func Test_Customer_HTTP_Create_DuplicateNameWithoutLastName(
	t *testing.T,
) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	_, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	request := api_models.CreateCustomerRequest{
		FirstName: firstName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
		rec,
		req,
	)

	if rec.Code == http.StatusCreated {
		t.Fatal("expected duplicate name error")
	}
}

func Test_Customer_HTTP_GetByID_InvalidUUID(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/customers/not-a-uuid",
		nil,
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
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

func Test_Customer_HTTP_Create_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
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

func Test_Customer_HTTP_Update_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPatch,
		"/customers/11111111-1111-1111-1111-111111111111",
		bytes.NewBufferString("{invalid json"),
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
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

func Test_Customer_HTTP_Create_EmptyFirstName(t *testing.T) {
	request := api_models.CreateCustomerRequest{}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/customers",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomersMainHandler(
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

func Test_Customer_HTTP_GetByID(t *testing.T) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodGet,
		"/customers/"+created.Id,
		nil,
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
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

	var response api_models.Customer

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.Id != created.Id {
		t.Fatal("invalid customer id")
	}

	if response.FirstName != firstName {
		t.Fatal("invalid first name")
	}
}

func Test_Customer_HTTP_GetByID_NotFound(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/customers/11111111-1111-1111-1111-111111111111",
		nil,
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusNotFound,
			rec.Code,
			rec.Body.String(),
		)
	}
}

func Test_Customer_HTTP_Update_NotFound(t *testing.T) {
	newName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	request := api_models.UpdateCustomerRequest{
		FirstName: &newName,
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/customers/11111111-1111-1111-1111-111111111111",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	customers.CustomerByIDMainHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected %d got %d body=%s",
			http.StatusNotFound,
			rec.Code,
			rec.Body.String(),
		)
	}
}
