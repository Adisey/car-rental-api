# Добавление справочника Brand

## Цель

Добавить справочник производителей автомобилей.

Особенности:

- справочник независим от `cars`
- связь с `cars` будет добавлена позже
- поля:
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

Добавить endpoint:

```yaml
/brands:
```

GET и POST.

Добавить endpoint:

```yaml
/brands/{id}:
```

GET и PATCH.

DELETE не добавлять.

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

# 3. Создать модель БД

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

# 4. Зарегистрировать модель в генераторе схемы

Файл:

```text
cmd/schema/main.go
```

Добавить:

```go
&db_models.Brand{},
```

в список моделей.

---

# 5. Сгенерировать актуальную схему

Запустить:

```bash
go run ./cmd/schema > db/schema/schema.sql
```

Проверить изменения в:

```text
db/schema/schema.sql
```

Убедиться, что появилась таблица:

```sql
CREATE TABLE "brands"
```

и ограничение:

```sql
CONSTRAINT "idx_brands_name" UNIQUE ("name")
```

---

# 6. Создать миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate diff add_brands_dictionary --env local
```

---

# 7. Проверить созданную миграцию

Atlas должен создать таблицу:

```sql
CREATE TABLE "brands"
```

с полями:

```sql
id
name
```

и уникальным ограничением:

```sql
CONSTRAINT "idx_brands_name" UNIQUE ("name")
```

Проверить миграцию перед применением.

---

# 8. Добавить начальные данные

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

Почему данные добавляются в миграцию:

- одинаковые данные во всех окружениях
- готовый справочник сразу после применения миграции
- фронт может использовать его без ручного наполнения

---

# 9. Применить миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate apply --env local
```

---

# 10. Проверить результат

Проверить:

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

# 11. Создать репозиторий

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

# 12. Создать валидацию

Файл:

```text
internal/validation/brand.go
```

Создать:

```go
ValidateCreateBrandRequest()
```

```go
ValidateUpdateBrandRequest()
```

Проверки:

```text
required
min_length
max_length
```

---

# 13. Создать сервисы

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

Правило POST:

Если бренд найден по имени:

```go
GetByNameRepository(...)
```

вернуть существующий бренд.

Новый бренд не создавать.

Правило PATCH:

Если найден другой бренд с таким же именем:

```json
{
  "errors": {
    "name": [
      {
        "name": "already_exists"
      }
    ]
  }
}
```

---

# 14. Создать хендлеры

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

# 15. Зарегистрировать роуты

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

# 16. Проверочный набор curl

## Получить список

```bash
curl http://localhost:8080/brands
```

---

## Получить бренд по id

```bash
curl http://localhost:8080/brands/<id>
```

---

## Создать бренд

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{
  "name":"Alfa Romeo"
}'
```

---

## Повторное создание

```bash
curl -X POST http://localhost:8080/brands \
-H "Content-Type: application/json" \
-d '{
  "name":"Alfa Romeo"
}'
```

Ожидаем:

```text
существующий объект
```

Новая запись не создаётся.

---

## Проверка уникальности через PATCH

Пусть существуют:

```text
Audi
BMW
```

Выполняем:

```bash
curl -X PATCH http://localhost:8080/brands/<bmw-id> \
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

## Успешный PATCH

```bash
curl -X PATCH http://localhost:8080/brands/<id> \
-H "Content-Type: application/json" \
-d '{
  "name":"Volkswagen Group"
}'
```

---

## Ошибка валидации

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

# Definition of Done

- OpenAPI обновлён
- Go API модели сгенерированы
- модель Brand создана
- schema.sql обновлён
- миграция создана
- миграция применена
- бренды заполнены начальными данными
- GET список работает
- GET по id работает
- POST идемпотентен
- PATCH работает
- дубликаты имени запрещены
- валидация работает
- все curl-проверки проходят
