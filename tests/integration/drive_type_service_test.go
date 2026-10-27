//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	drive_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/drive_type"
)

func Test_DriveType_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	createRequest := api_models.CreateCodeEntityRequest{
		Code: fmt.Sprintf(
			"TEST_DRIVE_TYPE_%d",
			time.Now().UnixNano(),
		),
	}

	created, err := drive_type.CreateDriveTypeService(
		ctx,
		createRequest,
	)
	if err != nil {
		t.Fatalf("ошибка создания привода: %v", err)
	}

	if created == nil {
		t.Fatal("созданный объект равен nil")
	}

	if created.Code != createRequest.Code {
		t.Fatalf(
			"ожидался код %q, получен %q",
			createRequest.Code,
			created.Code,
		)
	}

	items, err := drive_type.GetDriveTypesService(ctx)
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

	updatedCode := fmt.Sprintf(
		"UPDATED_DRIVE_TYPE_%d",
		time.Now().UnixNano(),
	)

	updateRequest := api_models.UpdateCodeEntityRequest{
		Code: &updatedCode,
	}

	updated, err := drive_type.UpdateDriveTypeService(
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

	actual, err := drive_type.GetDriveTypeByIDService(
		ctx,
		created.Id,
	)
	if err != nil {
		t.Fatalf("ошибка получения по id: %v", err)
	}

	if actual.Code != updatedCode {
		t.Fatalf(
			"ожидался код %q, получен %q",
			updatedCode,
			actual.Code,
		)
	}
}

func Test_DriveType_Service_Create_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	code := fmt.Sprintf(
		"DUPLICATE_DRIVE_TYPE_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateCodeEntityRequest{
		Code: code,
	}

	first, err := drive_type.CreateDriveTypeService(
		ctx,
		request,
	)
	if err != nil {
		t.Fatalf("ошибка первого создания: %v", err)
	}

	second, err := drive_type.CreateDriveTypeService(
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

func Test_DriveType_Service_Update_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	firstCode := fmt.Sprintf(
		"FIRST_%d",
		time.Now().UnixNano(),
	)

	secondCode := fmt.Sprintf(
		"SECOND_%d",
		time.Now().UnixNano(),
	)

	_, err := drive_type.CreateDriveTypeService(
		ctx,
		api_models.CreateCodeEntityRequest{
			Code: firstCode,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания первой записи: %v", err)
	}

	second, err := drive_type.CreateDriveTypeService(
		ctx,
		api_models.CreateCodeEntityRequest{
			Code: secondCode,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания второй записи: %v", err)
	}

	request := api_models.UpdateCodeEntityRequest{
		Code: &firstCode,
	}

	_, err = drive_type.UpdateDriveTypeService(
		ctx,
		second.Id,
		request,
	)

	if err == nil {
		t.Fatal("ожидалась ошибка уникальности, получен nil")
	}
}

func Test_DriveType_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := drive_type.GetDriveTypeByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("ожидалась ошибка для несуществующего id")
	}
}
