//go:build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Adisey/car-rental-api/internal/api_models"
	body_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
)

/*
Интеграционный тест полного жизненного цикла BodyType.

Сценарий:

1. Создаём новый тип кузова.
2. Получаем список типов кузова.
3. Проверяем наличие созданной записи.
4. Обновляем запись.
5. Получаем запись по ID.
6. Проверяем результат обновления.

Проверяем:

- PostgreSQL;
- миграции;
- repository;
- service;
- чтение и запись данных.
*/
func Test_BodyType_Service_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	createRequest := api_models.CreateCodeEntityRequest{
		Code: fmt.Sprintf(
			"TEST_BODY_TYPE_%d",
			time.Now().UnixNano(),
		),
	}

	created, err := body_type.CreateBodyTypeService(
		ctx,
		createRequest,
	)
	if err != nil {
		t.Fatalf("ошибка создания типа кузова: %v", err)
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

	items, err := body_type.GetBodyTypesService(ctx)
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
		"UPDATED_BODY_TYPE_%d",
		time.Now().UnixNano(),
	)

	updateRequest := api_models.UpdateCodeEntityRequest{
		Code: &updatedCode,
	}

	updated, err := body_type.UpdateBodyTypeService(
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

	actual, err := body_type.GetBodyTypeByIDService(
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

/*
Проверяем повторное создание записи с одинаковым кодом.

Текущая бизнес-логика сервиса не создаёт дубликат,
а возвращает уже существующую запись.
*/
func Test_BodyType_Service_Create_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	code := fmt.Sprintf(
		"DUPLICATE_BODY_TYPE_%d",
		time.Now().UnixNano(),
	)

	request := api_models.CreateCodeEntityRequest{
		Code: code,
	}

	first, err := body_type.CreateBodyTypeService(
		ctx,
		request,
	)
	if err != nil {
		t.Fatalf("ошибка первого создания: %v", err)
	}

	second, err := body_type.CreateBodyTypeService(
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

/*
Проверяем ограничение уникальности кода при обновлении.

Сценарий:

1. Создаём первую запись.
2. Создаём вторую запись.
3. Пытаемся обновить вторую запись кодом первой.
4. Ожидаем ошибку валидации.
*/

func Test_BodyType_Service_Update_DuplicateCode(t *testing.T) {
	ctx := context.Background()

	firstCode := fmt.Sprintf(
		"FIRST_%d",
		time.Now().UnixNano(),
	)

	secondCode := fmt.Sprintf(
		"SECOND_%d",
		time.Now().UnixNano(),
	)

	_, err := body_type.CreateBodyTypeService(
		ctx,
		api_models.CreateCodeEntityRequest{
			Code: firstCode,
		},
	)
	if err != nil {
		t.Fatalf("ошибка создания первой записи: %v", err)
	}

	second, err := body_type.CreateBodyTypeService(
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

	_, err = body_type.UpdateBodyTypeService(
		ctx,
		second.Id,
		request,
	)

	if err == nil {
		t.Fatal("ожидалась ошибка уникальности, получен nil")
	}
}

/*
Проверяем получение несуществующей записи.

Ожидаем ошибку репозитория/сервиса.
*/
func Test_BodyType_Service_GetByID_NotFound(t *testing.T) {
	ctx := context.Background()

	_, err := body_type.GetBodyTypeByIDService(
		ctx,
		"11111111-1111-1111-1111-111111111111",
	)

	if err == nil {
		t.Fatal("ожидалась ошибка для несуществующего id")
	}
}
