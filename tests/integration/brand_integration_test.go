//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
)

func Test_Brand_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	createRequest := api_models.CreateBrandRequest{
		Name: fmt.Sprintf(
			"TEST_BRAND_%d",
			time.Now().UnixNano(),
		),
	}

	created, err := brand.CreateBrandService(
		ctx,
		createRequest,
	)
	if err != nil {
		t.Fatalf("ошибка создания бренда: %v", err)
	}

	if created == nil {
		t.Fatal("созданный объект равен nil")
	}

	if created.Name != createRequest.Name {
		t.Fatalf(
			"ожидалось имя %q, получено %q",
			createRequest.Name,
			created.Name,
		)
	}

	items, err := brand.GetBrandsService(ctx)
	if err != nil {
		t.Fatalf("ошибка получения списка: %v", err)
	}

	found := false

	for _, item := range items {
		if item.Id == created.Id {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf(
			"созданная запись %s отсутствует в списке",
			created.Id,
		)
	}

	updatedName := fmt.Sprintf(
		"UPDATED_BRAND_%d",
		time.Now().UnixNano(),
	)

	updateRequest := api_models.UpdateBrandRequest{
		Name: &updatedName,
	}

	updated, err := brand.UpdateBrandService(
		ctx,
		created.Id,
		updateRequest,
	)
	if err != nil {
		t.Fatalf("ошибка обновления: %v", err)
	}

	if updated == nil {
		t.Fatal("обновлённый объект равен nil")
	}

	actual, err := brand.GetBrandByIDService(
		ctx,
		created.Id,
	)
	if err != nil {
		t.Fatalf("ошибка получения по id: %v", err)
	}

	if actual.Name != updatedName {
		t.Fatalf(
			"ожидалось имя %q, получено %q",
			updatedName,
			actual.Name,
		)
	}
}

func Test_Brand_Create_DuplicateName(t *testing.T) {
	ctx := context.Background()

	name := fmt.Sprintf(
		"DUPLICATE_BRAND_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateBrandRequest{
		Name: name,
	}

	first, err := brand.CreateBrandService(
		ctx,
		request,
	)
	if err != nil {
		t.Fatalf("ошибка первого создания: %v", err)
	}

	second, err := brand.CreateBrandService(
		ctx,
		request,
	)
	if err != nil {
		t.Fatalf("ошибка повторного создания: %v", err)
	}

	if first.Id != second.Id {
		t.Fatalf(
			"ожидался возврат существующей записи, получены разные ID: %s и %s",
			first.Id,
			second.Id,
		)
	}
}

func Test_Brand_Update_DuplicateName(t *testing.T) {
	ctx := context.Background()

	firstName := fmt.Sprintf(
		"FIRST_%d",
		time.Now().UnixNano(),
	)

	secondName := fmt.Sprintf(
		"SECOND_%d",
		time.Now().UnixNano(),
	)

	_, err := brand.CreateBrandService(
		ctx,
		api_models.CreateBrandRequest{
			Name: firstName,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания первой записи: %v", err)
	}

	second, err := brand.CreateBrandService(
		ctx,
		api_models.CreateBrandRequest{
			Name: secondName,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания второй записи: %v", err)
	}

	request := api_models.UpdateBrandRequest{
		Name: &firstName,
	}

	_, err = brand.UpdateBrandService(
		ctx,
		second.Id,
		request,
	)

	if err == nil {
		t.Fatal("ожидалась ошибка уникальности, получен nil")
	}
}

func Test_Brand_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := brand.GetBrandByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("ожидалась ошибка для несуществующего id")
	}
}