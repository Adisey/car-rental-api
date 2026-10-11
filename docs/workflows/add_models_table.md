# Добавление справочника Models

## Цель

Добавить справочник моделей автомобилей.

Примеры:

- Golf
- Passat
- X5
- A4

Особенности:

- модель принадлежит бренду
- модель может использоваться автомобилями
- удаление не реализуется
- используется связь:

```text
Brand
  ↓
Model
```

---

# 1. Обновить OpenAPI

Добавить схему:

```yaml
Model:
  type: object

  required:
    - id
    - name

  properties:
    id:
      type: string

    name:
      type: string
      minLength: 2
      maxLength: 100

    brand:
      allOf:
        - $ref: "#/components/schemas/Brand"
      nullable: true
```

---

Добавить схему:

```yaml
ModelBrand:
  type: object

  description: |
    Brand reference.

    Resolution order:

    1. id
    2. name

  properties:
    id:
      type: string

    name:
      type: string
```

---

Добавить схему:

```yaml
CreateModelRequest:
  type: object

  required:
    - name

  properties:
    name:
      type: string
      minLength: 2
      maxLength: 100

    brand:
      allOf:
        - $ref: "#/components/schemas/ModelBrand"
      nullable: true
```

---

Добавить схему:

```yaml
UpdateModelRequest:
  type: object

  properties:
    name:
      type: string
      minLength: 2
      maxLength: 100
      nullable: true

    brand:
      allOf:
        - $ref: "#/components/schemas/ModelBrand"
      nullable: true
```

---

Добавить endpoint:

```yaml
GET /models
GET /models/{id}
POST /models
PATCH /models/{id}
```

---

Для списка добавить фильтр:

```yaml
brandId:
  in: query
```

Пример:

```http
GET /models?brandId=<uuid>
```

---

# 2. Сгенерировать Go API модели

Запустить:

```bash
./cmd/openapi/generate.sh
```

Проверить появление:

```go
type Model struct
type CreateModelRequest struct
type UpdateModelRequest struct
type ModelBrand struct
```

---

# 3. Обновить db_model

Создать файл:

internal/db_models/model.go

````go
package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Model struct {
	bun.BaseModel `bun:"table:models"`

	ID      uuid.UUID  `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name    string     `bun:"type:text,notnull,unique:idx_models_name_brand"`
	BrandID *uuid.UUID `bun:"type:uuid"*

	Brand *Brand `bun:"rel:belongs-*o,join:brand_id=id"`
}

Обновить:

internal/db_models/car.go

Добавить:
```go
BrandID *uuid.UUID`bun:"type:uuid"`
ModelID *uuid.UUID `bun:"type:uuid"`

Bran* *Brand `bun:"rel:belongs-to,join:brand_id=id"`
Model *Model `bun:"rel:belongs-to,join:model_id=id"`
````

---

# 4. Сгененерировать актуальную схему

Зарегистрировать модель в:

```text
cmd/schema/main.go
```

Добавить:

```go
*db_models.Model{},
```

в список моделей.

---

Запустить:

```bash
 go run ./cmd/schema > db/schema/schema.sql
```

---

# 5. Проверить db/schema/schema.sql

Проверить, что новая таблица присутствует в схеме.

Пример:

```sql
CREATE TABLE "models"
```

---

Проверить наличие всех ожидаемых колонок.

Пример:

```sql
id
name
brand_id
```

---

Проверить наличие всех ожидаемых внешних ключей.

Пример:

```sql
FOREIGN KEY ("brand_id")
REFERENCES "brands" ("id")
```

---

Проверить наличие всех ожидаемых ограничений уникальности.

Пример:

```sql
CONSTRAINT "idx_models_name_brand"
UNIQUE ("name", "brand_id")
```

---

Проверить, что в schema.sql присутствуют все ограничения,
описанные в db_models.

Особое внимание уделять:

- UNIQUE
- FOREIGN KEY
- CHECK
- DEFAULT

---

Для справочников с уникальными строковыми значениями:

```text
Brand.name
Color.code
```

дополнительно проверить реализацию поиска без учёта регистра
на уровне Go-сервисов.

Пример:

```text
BMW
bmw
BmW
```

должны рассматриваться как один бренд.

```text
RED
red
Red
```

должны рассматриваться как один цвет.

# 6. Создать миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate diff add_models_dictionary --env local
```

---

# 7. Проверить миграцию

Убедиться в наличии:

```sql
CREATE *ABLE models
```

---

Проверить наличие:

```sql
FOREIGN KEY (brand_i*)
REFERENCES brands(id)
```

---

Проверить наличие ограничения:

```*ql
UNIQUE(name, brand_id)
```

---

Проверить отсутствие неожиданных:\_

```sql
DROP ...
```

операций.

---

# 8. Применить миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate apply --env local
```

---

# 9. Проверить результат в БД

Проверить:

```sql
SELECT *
FROM models;
```

---

Проверить внешний клю\*:

```sql
SELECT
	tc.constraint_na*e
FROM information_schema.table_co*straints tc
WHERE tc.table_name = *models'
  AND tc.constraint_type =*'FOREIGN KEY';
```

---

# 10. Создать Repository

Создать:

```text
internal/repositories/model_repository.go
```

Интерфейс должен содержать:

```go
GetAllRepository
GetByI*Repository
GetByNameRepository
Cre*teRepository
UpdateRepository
```

---

Реализация:

```text
internal/repositories/model_repository_bun.go
```

---

Поиск по имени выполняется без учёта регистра.

Использовать:

```sql
LOWER(name)
```

```go
package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	api_errors "github.com/Adisey/car-rental-api/internal/api_errors"
	"github.com/Adisey/car-rental-api/internal/db"
	"github.com/Adisey/car-rental-api/internal/db_models"
)

type BunModelRepository struct{}

func NewModelRepository() *BunModelRepository {
	return &BunModelRepository{}
}

func (r *BunModelRepository) GetAllRepository(
	ctx context.Context,
	brandID *uuid.UUID,
) ([]db_models.Model, error) {

	var models []db_models.Model

	query := db.BunDB.NewSelect().
		Model(&models).
		Relation("Brand")

	if brandID != nil {
		query = query.Where(
			"model.brand_id = ?",
			*brandID,
		)
	}

	err := query.Scan(ctx)

	return models, err
}

func (r *BunModelRepository) GetByIDRepository(
	ctx context.Context,
	id uuid.UUID,
) (*db_models.Model, error) {

	model := new(db_models.Model)

	err := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("model.id = ?", id).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *BunModelRepository) GetByNameRepository(
	ctx context.Context,
	name string,
	brandID *uuid.UUID,
) (*db_models.Model, error) {

	model := new(db_models.Model)

	query := db.BunDB.NewSelect().
		Model(model).
		Relation("Brand").
		Where("LOWER(model.name) = LOWER(?)", name)

	if brandID != nil {
		query = query.Where(
			"model.brand_id = ?",
			*brandID,
		)
	}

	err := query.Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, api_errors.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *BunModelRepository) CreateRepository(
	ctx context.Context,
	model *db_models.Model,
) error {

	_, err := db.BunDB.NewInsert().
		Model(model).
		Exec(ctx)

	return err
}

func (r *BunModelRepository) UpdateRepository(
	ctx context.Context,
	id uuid.UUID,
	request api_models.UpdateModelRequest,
	brandID *uuid.UUID,
	shouldUpdateBrand bool,
) (*db_models.Model, error) {

	model, err := r.GetByIDRepository(
		ctx,
		id,
	)

	if err != nil {
		return nil, err
	}

	if request.Name != nil {
		model.Name = *request.Name
	}

	if shouldUpdateBrand {
		model.BrandID = brandID
	}

	_, err = db.BunDB.NewUpdate().
		Model(model).
		WherePK().
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return model, nil
}
```

---

- 11. Создать Validation

Создать:

```text
internal/validation/model.go

```

---

Добавить:

```go
ModelRules.Name.MinLength = 2
ModelRules.Name.MaxLength = 100
```

---

Создать:

```go
ValidateCreateModelRequest()
ValidateUpdateModelRequest()
```

---

# 12. Создать Services

Создать:

```go
GetModelsService()
GetModelByIDService()
CreateModelService()
UpdateModelService()
```

---

Логика Brand должна работать по\*тем же правилам, что и Color.

---

Разрешение бренда:

```json
{
  "*rand": {
    "id": "..."
  }
}
```

использовать бренд по id.

---

```json
{
  "brand": {
    "name"* "Volkswagen"
  }
}
```

↓

## искать бренд по имени без учёта регистра.

Если бренд найден:

```text
использовать найденный бренд
```

---

Если бренд не найден:

```text
создать бренд
использовать новый бренд
```

---

Приоритет:

```text
id
↓
name
```

---

# 13. Обновить Service слой

Проверить маппинг:

```text
db_models -> api_models

```

---

Для ответа:

```json
{
  "id":*"...",
  "name": "Golf",
  "brand"* {
    "id": "...",
    "name": "Volkswagen"
  }
}
```

---

# 14. Сгeнерировать Frontend Types

Во фронтенде:

````bash
npm run generate
``

---

# 15. Обновить Frontend

Добавить поддержку:

```http
GET /models?brandId=<uuid>
````

---

# 16. Набор curl для проверки

## Получить список моделей

```bash
curl http://localhost:8080/models
```

---

## Получить список моделей бренда

```bash
curl "http://localhost:8080/models?brandId=59c200b2-965f-4674-95a0-1a56d96d703b"
```

Ожидаем только модели указанного бренда.

---

## Создать модель через существующий brand.id

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf",
  "brand":{
    "id":"<brand-id>"
  }
}'
```

---

## Создать модель через существующий brand.name

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Passat",
  "brand":{
    "name":"Volkswagen"
  }
}'
```

---

## Создать модель через новый brand.name

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"X6",
  "brand":{
    "name":"BMW"
  }
}'
```

Ожидаем:

- создание бренда (если его не было);
- создание модели;
- привязку модели к бренду.

---

## Проверка поиска бренда без учёта регистра

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Tiguan",
  "brand":{
    "name":"volkswagen"
  }
}'
```

Ожидаем использование существующего бренда Volkswagen.

---

## Повторный POST модели

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf",
  "brand":{
    "name":"Volkswagen"
  }
}'
```

Ожидаем возврат существующей модели без создания новой записи.

---

## Проверка одинаковых моделей у разных брендов

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf",
  "brand":{
    "name":"BMW"
  }
}'
```

Ожидаем успешное создание новой модели.

---

## Получить модель по id

```bash
curl http://localhost:8080/models/d2d68de4-12d7-4241-ab71-d2dccc358e51
```

Ожидаем:

```json
{
  "id": "...",
  "name": "Golf",
  "brand": {
    "id": "...",
    "name": "Volkswagen"
  }
}
```

---

## PATCH имени модели

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf VII"
}'
```

---

## PATCH через существующий brand.id

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "id":"59c200b2-965f-4674-95a0-1a56d96d703b"
  }
}'
```

---

## PATCH через существующий brand.name

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "name":"Audi"
  }
}'
```

---

## PATCH через новый brand.name

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "name":"NewBrand"
  }
}'
```

Ожидаем:

- создание нового бренда;
- привязку модели к новому бренду.

---

## Очистка бренда

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff  \
-H "Content-Type: application/json" \
-d '{
  "brand": {}
}'
```

Ожидаем:

```text
brand_id становится NULL
```

---

## Несуществующий brand.id

```bash
curl -X PATCH http://localhost:8080/models/43179486-c091-4ab5-9e5a-1f4e4f08adff \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "id":"00000000-0000-0000-0000-000000000000"
  }
}'
```

Ожидаем:

```json
{
  "errors": {
    "brand": [
      {
        "code": "not_found"
      }
    ]
  }
}
```

---

## Создать модель через существующий brand.id

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf",
  "brand":{
    "id":"<brand-id>"
  }
}'
```

---

## Создать модель через существующий brand.name

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"Passat",
  "brand":{
    "name":"Volkswagen"
  }
}'
```

---

## Создать модель через новый brand.name

```bash
curl -X POST http://localhost:8080/models \
-H "Content-Type: application/json" \
-d '{
  "name":"X6",
  "brand":{
    "name":"BMW"
  }
}'
```

Ожидаем:

- создание бренда (если его не было);
- создание модели;
- привязку модели к бренду.

---

## Проверка поиска бренда без учёта регистра

````bash
curl -X POST http://localhost:

И создание Машины

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Членовоз 2",
  "description":"Resolve all references",
  "color":{
    "code":"dark-blue"
  },
  "brand":{
    "name":"Aurus"
  },
  "model":{
    "name":"Long",
  }
}'
````

```bash
curl -X GET http://localhost:8080/cars
```
