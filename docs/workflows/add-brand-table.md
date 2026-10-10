# Добавление справочника Brand + привязка к Car

## Цель

Добавить справочник производителей автомобилей и сразу подготовить полноценную связь с таблицей `cars`.

Особенности:

- справочник независим от `cars`
- используется связь через `brand_id`
- поля справочника:
  - `id`
  - `name`
- нет полей:
  - `created_at`
  - `updated_at`
  - `deleted_at`
- удаление не реализуется
- при POST, если бренд с таким `name` уже существует:
  - ошибка не возвращается
  - возвращается существующая запись
- при PATCH нельзя изменить бренд на имя уже существующего бренда

---

# 1. Обновить OpenAPI

Файл:

```text
openapi/api.yaml
```

Добавить схему:

```yaml
Brand:
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
```

Добавить:

```yaml
CreateBrandRequest:
  type: object

  required:
    - name

  properties:
    name:
      type: string
      minLength: 2
      maxLength: 100
```

Добавить:

```yaml
UpdateBrandRequest:
  type: object

  properties:
    name:
      type: string
      minLength: 2
      maxLength: 100
      nullable: true
```

---

# 2. Сгенерировать Go API модели

Запустить:

```bash
./cmd/openapi/generate.sh
```

Проверить появление:

```go
type Brand struct
```

```go
type CreateBrandRequest struct
```

```go
type UpdateBrandRequest struct
```

в:

```text
internal/api_models/types.gen.go
```

---

# 3. Создать модель БД Brand

Файл:

```text
internal/db_models/brand.go
```

```go
package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Brand struct {
	bun.BaseModel `bun:"table:brands"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:"type:text,notnull,unique:idx_brands_name"`
}
```

---

# 4. Добавить brand_id в Car

Файл:

```text
internal/db_models/car.go
```

Добавить:

```go
BrandID *uuid.UUID `bun:"type:uuid"`

Brand Brand `bun:"rel:belongs-to,join:brand_id=id"`
```

Пример:

```go
ty*e Car struct {
	bun.BaseModel `bun*"table:cars"`

	ID          uuid.UUID  `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name        string     `bun:"type:text,notnull"`
	Description *string    `bun:"type:text"`
	ColorID     *uuid.UUID `bun:"type:uuid"`
	Color       *Color     `bun:"rel:belongs-to,join:color_id=id"`
	BrandID     *uuid.UUID `bun:"type:uuid"`
	Brand       Brand      `bun:"rel:belongs-to,join:brand_id=id"`
	CreatedAt   time.Time  `bun:",notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `bun:",notnull,default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `bun:"type:timestamptz"`
}
``*

---

# 5. Зарегистрировать модел* Brand в генераторе схем

Файл:

`*`text
cmd/schema/main.go
```

Доба\*ить:

````go
&db_models.Brand{},
``*

в список моделей.

---

# 6. Сгенерировать актуальную схему

Запуст*ть:

```bash
go run ./cmd/schema >*db/schema/schema.sql
````

Проверит\*:

````sql
CREATE TABLE "brands"
``*

и

```sql
CONSTRAINT "idx_brands*name" UNIQUE ("name")
````

Провери*ь, что в таблице `cars` появились:*

```sql
brand_id uuid
```

и

```s*l
FOREIGN KEY ("brand_id")
REFEREN*ES "brands" ("id")
```

---

# 7. Создать миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate diff add_brands_dictionary --env local
```

---

# 8. Проверить созданную миграцию

Atlas должен создать:

\*``sql
CREATE TABLE brands

````

и:
*```sql
ALTER TABLE cars
ADD COLUMN*brand_id uuid NULL
````

а также внешний ключ:

```sql
FOREIGN KEY (br*nd_id)
REFERENCES brands(id)
```

Проверить миграцию перед применением.

---

# 9. Добавить начальные да\*ные

В конец миграции добавить:

```sql
INSERT INTO brands (name)
VALUES
('Audi'),
('BMW'),
('Citroen'),
('Cupra'),
('Dacia'),
('Fiat'),
('Ford'),
('Honda'),
('Hyundai'),
('Jeep'),
('Kia'),
('Land Rover'),
('Mazda'),
('Mercedes-Benz'),
('Mini'),
('Mitsubishi'),
('Nissan'),
('Opel'),
('Peugeot'),
('Porsche'),
('Renault'),
('Seat'),
('Skoda'),
('Suzuki'),
('Tesla'),
('Toyota'),
('Volkswagen'),
('Volvo');
```

## Если осознанно менял миграции вручную, то нужно пересчитать контрольные суммы:

./cmd/schema/atlas.sh migrate hash --env local

---

# 10. Применить миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate apply --env local
```

---

# 11. Проверить результат

```sql
SELECT *
FROM brands
ORDER BY name;
```

Проверить количество записей:

```sql
SELECT COUNT(*)
FROM brands;
```

Ожидаем:

```text
28
```

---

Проверить таблицу cars:

```sql
SELECT
	column_name,
	is_nullable
FROM information_schema.columns
WHERE table_name = 'cars'
AND column_name = 'brand_id';
```

Ожидаем:

```text
brand_id | YES
```

---

# 12. Создать репозиторий

Файлы:

```text
internal/repositories/brand_repository.go
internal/repositories/bun_brand_repository.go
```

Добавить методы:

```go
GetAllRepository()
GetByIDRepository()
GetByNameRepository()
CreateRepository()
UpdateRepository()
```

---

# 13. Создать валидацию

Файл:

```text
internal/validation/brand.go
```

Создать:

```go
ValidateCreateBrandRequest()
ValidateUpdateBrandRequest()
```

Проверки:

```text
required
min_length
max_length
```

---

# 14. Создать сервисы

Файл:

```text
internal/services/brand_service.go
```

Создать:

```go
GetBrandsService()
GetBrandByIDService()
CreateBrandService()
UpdateBrandService()
```

Особенности:

POST:

Если бренд найден по имени:

```go
GetByNameRepository(...)
```

вернуть существующую запись.

Новый бренд не создавать.

---

PATCH:

Если найден другой бренд с таким же именем:

```json
{
  "errors": {
    "name": [
      {
        "code": "already_exists"
      }
    ]
  }
}
```

---

# 15. Создать хендлеры

Файл:

```text
internal/handlers/brand_handler.go
```

Создать:

```go
getBrandsHandler()
getBrandByIDHandler()
createBrandHandler()
updateBrandHandler()
```

Создать:

```go
BrandsMainHandler()
BrandByIDMainHandler()
```

---

# 16. Зарегистрировать роуты

Файл:

```text
cmd/server/main.go
```

Добавить:

```go
http.HandleFunc(
	"/brands",
	handlers.BrandsMainHandler,
)

http.HandleFunc(
	"/brands/",
	handlers.BrandByIDMainHandler,
)
```

---

# 17. Добавить Brand в ответы Car

Обновить OpenAPI.

Схема Car должна содержать:

```yaml
brand:
  allOf:
    - $ref: "#/components/schemas/Brand"
  nullable: true
```

---

После генерации обновить:

```text
GetCarsService()
GetCarByIDService()
```

по аналогии с Color.

---

Ожидаемый ответ:

```json
{
  "id": "...",
  "name": "Golf VII",
  "brand": {
    "id": "...",
    "name": "Volkswagen"
  }
}
```

---

# 18. Проверочный набор curl для Brand

## Получить список брендов

```bash
curl http://localhost:8080/brands
```

Ожидаем список брендов.

---

## Получить бренд по id

```bash
curl http://localhost:8080/brands/<BRAND_ID>
```

Ожидаем:

```json
{
  "id": "...",
  "name": "Audi"
}
```

---

## Создать новый бренд

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{
  "name":"Alfa Romeo"
}'
```

Ожидаем:

```json
{
  "id": "...",
  "name": "Alfa Romeo"
}
```

---

## Повторное создание существующего бренда

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{
  "name":"Alfa Romeo"
}'
```

Ожидаем:

```json
{
  "id": "...",
  "name": "Alfa Romeo"
}
```

Новая запись не создаётся.

---

## Проверка минимальной длины

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{
  "name":"A"
}'
```

Ожидаем:

```json
{
  "errors": {
    "name": [
      {
        "code": "min_length",
        "params": {
          "min": 2
        }
      }
    ]
  }
}
```

---

## Проверка обязательности поля

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{}'
```

Ожидаем:

```json
{
  "errors": {
    "name": [
      {
        "code": "required"
      }
    ]
  }
}
```

---

## Получить список брендов и выбрать id

```bash
curl http://localhost:8080/brands
```

Запомнить любой id.

---

## Переименовать бренд

```bash
curl -X PATCH http://localhost:8080/brands/<BRAND_ID> \
-H "Content-Type: application/json" \
-d '{
  "name":"Volkswagen Group"
}'
```

Ожидаем:

```json
{
  "id": "...",
  "name": "Volkswagen Group"
}
```

---

## Проверка дубликата через PATCH

Пусть существуют:

```text
Audi
BMW
```

Получаем id BMW и выполняем:

```bash
curl -X PATCH http://localhost:8080/brands/<BMW_ID> \
-H "Content-Type: application/json" \
-d '{
  "name":"Audi"
}'
```

Ожидаем:

```json
{
  "errors": {
    "name": [
      {
        "code": "already_exists"
      }
    ]
  }
}
```

---

# Проверка Car + Brand

## Создать машину без бренда

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Tesla Model Y"
}'
```

Ожидаем:

```json
{
  "brand": null
}
```

или отсутствие поля brand.

---

## Создать машину через существующий Brand ID

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Golf VII",
  "brand":{
    "id":"<BRAND_ID>"
  }
}'
```

Ожидаем:

```json
{
  "brand": {
    "id": "<BRAND_ID>",
    "name": "Volkswagen"
  }
}
```

---

## Создать машину через Brand Name

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Passat B8",
  "brand":{
    "name":"Volkswagen"
  }
}'
```

Ожидаем использование существующего бренда.

---

## Создать машину через новый Brand Name

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Giulia",
  "brand":{
    "name":"Alfa Romeo"
  }
}'
```

Если бренд отсутствует:

```text
создаётся Brand
↓
создаётся Car
↓
назначается brand_id
```

---

## Создать машину через несуществующий Brand ID

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Test",
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

## PATCH машины без изменения бренда

```bash
curl -X PATCH http://localhost:8080/cars/<CAR_ID> \
-H "Content-Type: application/json" \
-d '{
  "name":"Updated name"
}'
```

Ожидаем:

```text
brand не меняется
```

---

## Очистить бренд

```bash
curl -X PATCH http://localhost:8080/cars/<CAR_ID> \
-H "Content-Type: application/json" \
-d '{
  "brand": {}
}'
```

Ожидаем:

```json
{
  "brand": null
}
```

---

## Назначить бренд через id

```bash
curl -X PATCH http://localhost:8080/cars/<CAR_ID> \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "id":"<BRAND_ID>"
  }
}'
```

Ожидаем:

```json
{
  "brand": {
    "id": "<BRAND_ID>",
    "name": "Volkswagen"
  }
}
```

---

## Назначить бренд через name

```bash
curl -X PATCH http://localhost:8080/cars/<CAR_ID> \
-H "Content-Type: application/json" \
-d '{
  "brand":{
    "name":"Audi"
  }
}'
```

Ожидаем:

```json
{
  "brand": {
    "name": "Audi"
  }
}
```

---

## Проверить Car после всех изменений

```bash
curl http://localhost:8080/cars/<CAR_ID>
```

Ожидаем:

```json
{
  "id": "...",
  "name": "...",
  "brand": {
    "id": "...",
    "name": "..."
  }
}
```
