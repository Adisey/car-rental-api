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
Интеграционный тест справочника типов кузова.

Сценарий:

1. Создаём новый тип кузова.
2. Получаем список типов кузова.
2.1 Проверяем, что созданная запись присутствует в списке.
3. Обновляем созданную запись.
3.1 Получаем обновлённую запись по ID.
3.2 Проверяем корректность обновления.

Что проверяет тест:

- работу PostgreSQL;
- работу миграций;
- работу репозитория;
- работу сервисного слоя;
- корректность чтения и записи данных.
*/
func Test_BodyType_Create_GetAll_Update_GetByID(t *testing.T) {
	ctx := context.Background()

	//------------------------------------------------------------------
	// 1. Создание
	//------------------------------------------------------------------

	createRequest := api_models.CreateCodeEntityRequest{
		Code: fmt.Sprintf("TEST_BODY_TYPE_%d", time.Now().UnixNano()),
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

	//------------------------------------------------------------------
	// 2. Получение списка
	//------------------------------------------------------------------

	items, err := body_type.GetBodyTypesService(ctx)
	if err != nil {
		t.Fatalf("ошибка получения списка: %v", err)
	}

	//------------------------------------------------------------------
	// 2.1 Поиск созданной записи в списке
	//------------------------------------------------------------------

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

	//------------------------------------------------------------------
	// 3. Обновление
	//------------------------------------------------------------------

	updatedCode := "UPDATED_BODY_TYPE"

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

	//------------------------------------------------------------------
	// 3.1 Получение по ID
	//------------------------------------------------------------------

	actual, err := body_type.GetBodyTypeByIDService(
		ctx,
		created.Id,
	)
	if err != nil {
		t.Fatalf("ошибка получения по id: %v", err)
	}

	//------------------------------------------------------------------
	// 3.2 Проверка результата обновления
	//------------------------------------------------------------------

	if actual.Code != updatedCode {
		t.Fatalf(
			"ожидался код %q, получен %q",
			updatedCode,
			actual.Code,
		)
	}
}
