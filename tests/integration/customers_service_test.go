//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	customers "github.com/Adisey/car-rental-api/internal/customers"
)

func Test_Customer_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

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
			FirstName: "Ivan",
			LastName:  &lastName,
			Phone:     &phone,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created == nil {
		t.Fatal("created customer is nil")
	}

	if created.Id == "" {
		t.Fatal("customer id is empty")
	}

	list, err := customers.GetCustomersService(ctx)

	if err != nil {
		t.Fatal(err)
	}

	found := false

	for _, item := range list {
		if item.Id == created.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("customer not found in list")
	}

	updatedFirstName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	updated, err := customers.UpdateCustomerService(
		ctx,
		created.Id,
		api_models.UpdateCustomerRequest{
			FirstName: &updatedFirstName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.FirstName != updatedFirstName {
		t.Fatal("first name not updated")
	}

	actual, err := customers.GetCustomerByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.FirstName != updatedFirstName {
		t.Fatal("incorrect first name after update")
	}

	if actual.LastName == nil {
		t.Fatal("last name disappeared")
	}

	if *actual.LastName != lastName {
		t.Fatal("last name changed unexpectedly")
	}

	if actual.Phone == nil {
		t.Fatal("phone disappeared")
	}

	if *actual.Phone != phone {
		t.Fatal("phone changed unexpectedly")
	}

	if actual.Email == nil {
		t.Fatal("email disappeared")
	}

	if *actual.Email != email {
		t.Fatal("email changed unexpectedly")
	}
}

func Test_Customer_Service_Create_DuplicateNameWithoutLastName(t *testing.T) {
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

	_, err = customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
		},
	)

	if err == nil {
		t.Fatal("expected duplicate name error")
	}
}

func Test_Customer_Service_Create_SameFirstNameDifferentLastName(t *testing.T) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)

	lastName1 := fmt.Sprintf(
		"Ivanov_%d",
		time.Now().UnixNano(),
	)

	lastName2 := fmt.Sprintf(
		"Petrov_%d",
		time.Now().UnixNano(),
	)

	_, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			LastName:  &lastName1,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			LastName:  &lastName2,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Customer_Service_Create_DuplicatePhone(t *testing.T) {
	ctx := context.Background()

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	_, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: "Ivan",
			Phone:     &phone,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: "Petr",
			Phone:     &phone,
		},
	)

	if err == nil {
		t.Fatal("expected duplicate phone error")
	}
}

func Test_Customer_Service_Create_DuplicateEmail(t *testing.T) {
	ctx := context.Background()

	email := fmt.Sprintf(
		"test_%d@test.com",
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

	_, err = customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: "Petr",
			Email:     &email,
		},
	)

	if err == nil {
		t.Fatal("expected duplicate email error")
	}
}

func Test_Customer_Service_Update_FirstNameOnly_KeepsOtherFields(
	t *testing.T,
) {
	ctx := context.Background()

	lastName := fmt.Sprintf(
		"Ivanov_%d",
		time.Now().UnixNano(),
	)

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	email := fmt.Sprintf(
		"test_%d@test.com",
		time.Now().UnixNano(),
	)

	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: "Ivan",
			LastName:  &lastName,
			Phone:     &phone,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newFirstName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	_, err = customers.UpdateCustomerService(
		ctx,
		created.Id,
		api_models.UpdateCustomerRequest{
			FirstName: &newFirstName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	actual, err := customers.GetCustomerByIDService(
		ctx,
		created.Id,
	)

	if err != nil {
		t.Fatal(err)
	}

	if actual.FirstName != newFirstName {
		t.Fatal("first name not updated")
	}

	if actual.LastName == nil {
		t.Fatal("last name disappeared")
	}

	if *actual.LastName != lastName {
		t.Fatal("last name changed unexpectedly")
	}

	if actual.Phone == nil {
		t.Fatal("phone disappeared")
	}

	if *actual.Phone != phone {
		t.Fatal("phone changed unexpectedly")
	}

	if actual.Email == nil {
		t.Fatal("email disappeared")
	}

	if *actual.Email != email {
		t.Fatal("email changed unexpectedly")
	}
}

func Test_Customer_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := customers.GetCustomerByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("expected error")
	}
}

func Test_Customer_Service_Create_ResponseContainsFields(
	t *testing.T,
) {
	ctx := context.Background()

	lastName := fmt.Sprintf(
		"Ivanov_%d",
		time.Now().UnixNano(),
	)

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)

	email := fmt.Sprintf(
		"test_%d@test.com",
		time.Now().UnixNano(),
	)

	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: "Ivan",
			LastName:  &lastName,
			Phone:     &phone,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if created == nil {
		t.Fatal("created customer is nil")
	}

	if created.Id == "" {
		t.Fatal("id is empty")
	}

	if created.FirstName != "Ivan" {
		t.Fatal("invalid first_name")
	}

	if created.LastName == nil {
		t.Fatal("last_name is nil")
	}

	if *created.LastName != lastName {
		t.Fatal("invalid last_name")
	}

	if created.Phone == nil {
		t.Fatal("phone is nil")
	}

	if *created.Phone != phone {
		t.Fatal("invalid phone")
	}

	if created.Email == nil {
		t.Fatal("email is nil")
	}

	if *created.Email != email {
		t.Fatal("invalid email")
	}
}

func Test_Customer_Service_Update_ResponseContainsFields(
	t *testing.T,
) {
	ctx := context.Background()

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)
	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)
	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			Phone:     &phone,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	updated, err := customers.UpdateCustomerService(
		ctx,
		created.Id,
		api_models.UpdateCustomerRequest{
			FirstName: &newName,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if updated.Id != created.Id {
		t.Fatal("id changed")
	}

	if updated.FirstName != newName {
		t.Fatal("first_name not updated")
	}

	if updated.Phone == nil {
		t.Fatal("phone disappeared")
	}

	if *updated.Phone != phone {
		t.Fatal("phone changed unexpectedly")
	}
}

func Test_Customer_Service_Update_SameCustomer_PhoneDoesNotConflict(
	t *testing.T,
) {
	ctx := context.Background()

	phone := fmt.Sprintf(
		"+371%d",
		time.Now().UnixNano(),
	)
	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)
	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
			Phone:     &phone,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	newName := fmt.Sprintf(
		"Petr_%d",
		time.Now().UnixNano(),
	)

	_, err = customers.UpdateCustomerService(
		ctx,
		created.Id,
		api_models.UpdateCustomerRequest{
			FirstName: &newName,
			Phone:     &phone,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_Customer_Service_Update_SameCustomer_EmailDoesNotConflict(
	t *testing.T,
) {
	ctx := context.Background()

	email := fmt.Sprintf(
		"test_%d@test.com",
		time.Now().UnixNano(),
	)
	firstName := fmt.Sprintf(
		"Ivan_%d",
		time.Now().UnixNano(),
	)
	created, err := customers.CreateCustomerService(
		ctx,
		api_models.CreateCustomerRequest{
			FirstName: firstName,
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

	_, err = customers.UpdateCustomerService(
		ctx,
		created.Id,
		api_models.UpdateCustomerRequest{
			FirstName: &newName,
			Email:     &email,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}
