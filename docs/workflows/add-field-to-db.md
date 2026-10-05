# Добавление нового поля в сущность

Пример: добавляем поля created_at, updated_at и deleted_at в сущность Car.

## 1. Обновить db_model

Файл:

internal/db_models/car.go

Добавить новые поля:

    CreatedAt time.Time  `bun:",notnull"`
    UpdatedAt time.Time  `bun:",notnull"`
    DeletedAt *time.Time `bun:"type:timestamptz"`

## 2. Сгенерировать актуальную схему

    go run ./cmd/schema > db/schema/schema.sql

## 3. Проверить изменения в schema.sql

Убедиться, что новая схема содержит ожидаемые колонки.

Пример:

    created_at timestamptz not null
    updated_at timestamptz not null
    deleted_at timestamptz null

Обязательно проверить:

DEFAULT gen_random_uuid()
DEFAULT CURRENT_TIMESTAMP
другие вычисляемые значения

## 4. Создать миграцию

Придумать осмысленное имя миграции.

Пример:

    ./cmd/schema/atlas.sh migrate diff add_car_timestamps --env local

## 5. Проверить созданную миграцию

Открыть файл в:

    db/migrations/

Проверить:

- типы колонок;
- nullable / not null;
- значения по умолчанию;
- индексы;
- отсутствие лишних изменений.

Нельзя применять миграцию без просмотра её содержимого.

## 6. Применить миграцию

    ./cmd/schema/atlas.sh migrate apply --env local

## 7. Проверить результат в базе данных

Проверить:

    SELECT *
    FROM cars
    LIMIT 1;

или через DBeaver.

Убедиться, что новые поля появились.

## 8. Обновить репозитории

Если новые поля участвуют в:

- создании;
- обновлении;
- фильтрации;
- сортировке;

обновить соответствующие Repository методы.

## 9. Обновить Service слой

Проверить маппинг:

    db_models -> api_models

пример:  
 internal/services/car_service.go

Если поле должно возвращаться клиенту, добавить его в преобразование моделей.

## 10. Обновить OpenAPI

Если поле должно быть доступно по API:

Обновить:

- schemas;
- request models;
- response models.

пример:
openapi/api.yaml

## 11. Перегенерировать API модели

Запустить принятую в проекте команду генерации типов.

пример:
"sync-api": "openapi-typescript ../car-rental-api/openapi/api.yaml -o src/types/api.ts",
"generate": "npm run sync-api && npm run lint"

## 12. Обновить Frontend

Подтянуть новые типы и использовать новые поля, если они участвуют в UI.

## 13. Выполнить ручную проверку

Проверить:

- GET список;
- GET карточку;
- создание;
- обновление;

если изменение затрагивает эти сценарии.

## 14. Коммит

Убедиться, что в коммит попали:

- db_models;
- schema.sql;
- migration.sql;
- repository изменения;
- service изменения;
- OpenAPI изменения;
- сгенерированные типы.

Сделать коммит только после успешного применения миграции и ручной проверки.
