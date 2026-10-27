//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
)

func Test_Color_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	createRequest := api_models.CreateColorRequest{
		Code: fmt.Sprintf(
			"TEST_COLOR_%d",
			time.Now().UnixNano(),
		),
	}

	created, err := color.CreateColorService(
		ctx,
		createRequest,
	)
	if err != nil {
		t.Fatalf("ошибка создания цвета: %v", err)
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

	if created.Hex != nil {
		t.Fatalf(
			"ожидался nil hex, получен %v",
			*created.Hex,
		)
	}

	items, err := color.GetColorsService(ctx)
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
		"UPDATED_COLOR_%d",
		time.Now().UnixNano(),
	)

	updatedHex := "#00FF00"

	updateRequest := api_models.UpdateColorRequest{
		Code: &updatedCode,
		Hex:  &updatedHex,
	}

	updated, err := color.UpdateColorService(
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

	actual, err := color.GetColorByIDService(
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

	if actual.Hex == nil {
		t.Fatal("hex отсутствует после обновления")
	}

	if *actual.Hex != updatedHex {
		t.Fatalf(
			"ожидался hex %q, получен %q",
			updatedHex,
			*actual.Hex,
		)
	}
}

func Test_Color_Service_Create_WithHex(t *testing.T) {
	ctx := context.Background()

	hex := "#FF0000"

	created, err := color.CreateColorService(
		ctx,
		api_models.CreateColorRequest{
			Code: fmt.Sprintf(
				"COLOR_%d",
				time.Now().UnixNano(),
			),
			Hex: &hex,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания цвета: %v", err)
	}

	if created.Hex == nil {
		t.Fatal("hex не сохранён")
	}

	if *created.Hex != hex {
		t.Fatalf(
			"ожидался hex %q, получен %q",
			hex,
			*created.Hex,
		)
	}
}

func Test_Color_Service_Create_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	code := fmt.Sprintf(
		"DUPLICATE_COLOR_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateColorRequest{
		Code: code,
	}

	first, err := color.CreateColorService(
		ctx,
		request,
	)
	if err != nil {
		t.Fatalf("ошибка первого создания: %v", err)
	}

	second, err := color.CreateColorService(
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

func Test_Color_Service_Update_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	firstCode := fmt.Sprintf(
		"FIRST_%d",
		time.Now().UnixNano(),
	)

	secondCode := fmt.Sprintf(
		"SECOND_%d",
		time.Now().UnixNano(),
	)

	_, err := color.CreateColorService(
		ctx,
		api_models.CreateColorRequest{
			Code: firstCode,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания первой записи: %v", err)
	}

	second, err := color.CreateColorService(
		ctx,
		api_models.CreateColorRequest{
			Code: secondCode,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания второй записи: %v", err)
	}

	request := api_models.UpdateColorRequest{
		Code: &firstCode,
	}

	_, err = color.UpdateColorService(
		ctx,
		second.Id,
		request,
	)

	if err == nil {
		t.Fatal("ожидалась ошибка уникальности, получен nil")
	}
}

func Test_Color_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := color.GetColorByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("ожидалась ошибка для несуществующего id")
	}
}
