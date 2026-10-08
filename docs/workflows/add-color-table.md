# Инструкция по добавлению справочника Colors

## Цель

Добавить справочник цветов автомобилей.

Особенности:

- справочник независим от `cars`
- связь с `cars` будет добавлена позже
- поля:
  - `id`
  - `code`
- нет полей:
  - `created_at`
  - `updated_at`
  - `deleted_at`
- удаление не реализуется
- при POST, если цвет с таким `code` уже существует:
  - ошибка не возвращается
  - возвращается существующая запись

---

# 1. OpenAPI

## Схема Color

Добавить в:

```yaml
components:
  schemas:
```

```yaml
Color:
  type: object

  required:
    - id
    - code

  properties:
    id:
      type: string

    code:
      type: string
      minLength: 2
      maxLength: 50
```

---

## Схема CreateColorRequest

```yaml
CreateColorRequest:
  type: object

  required:
    - code

  properties:
    code:
      type: string
      minLength: 2
      maxLength: 50
```

---

## Схема UpdateColorRequest

PATCH-подход как у Cars.

```yaml
UpdateColorRequest:
  type: object

  properties:
    code:
      type: string
      nullable: true
      minLength: 2
      maxLength: 50
```

---

## Endpoint GET /colors

```yaml
/colors:
  get:
    operationId: getColors

    responses:
      "200":
        description: Colors list

        content:
          application/json:
            schema:
              type: array

              items:
                $ref: "#/components/schemas/Color"
```

---

## Endpoint GET /colors/{id}

```yaml
/colors/{id}:
  get:
    operationId: getColorById

    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: string

    responses:
      "200":
        description: Color

        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Color"

      "404":
        description: Color not found
```

---

## Endpoint POST /colors

```yaml
/colors:
  post:
    operationId: createColor

    requestBody:
      required: true

      content:
        application/json:
          schema:
            $ref: "#/components/schemas/CreateColorRequest"

    responses:
      "201":
        description: Color created

        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Color"
```

---

## Endpoint PATCH /colors/{id}

```yaml
/colors/{id}:
  patch:
    operationId: updateColor

    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: string

    requestBody:
      required: true

      content:
        application/json:
          schema:
            $ref: "#/components/schemas/UpdateColorRequest"

    responses:
      "200":
        description: Updated color

        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Color"

      "404":
        description: Color not found
```

---

# 1.1 Сгенерировать Go API модели

После изменения:

```text
openapi/api.yaml
```

необходимо пересоздать:

```text
internal/api_models/types.gen.go
```

Запустить:

```bash
./cmd/openapi/generate.sh
```

Либо напрямую:

```bash
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest \
  -generate types \
  -package api_models \
  -o internal/api_models/types.gen.go \
  openapi/api.yaml
```

Проверить, что в:

```text
internal/api_models/types.gen.go
```

появились модели:

```go
type Color struct
```

```go
type CreateColorRequest struct
```

```go
type UpdateColorRequest struct
```

Проверить, что проект собирается:

```bash
go build ./...
```

``

# 2. Генерация моделей

После изменения OpenAPI:

```bash
./cmd/openapi/generate.sh
```

на фронте делаем импорт

```package.json
"sync-api": "openapi-typescript ../car-rental-api/openapi/api.yaml -o src/types/api.ts",
"generate": "npm run sync-api && npm run lint"
```

---

# 3. Модель БД

Файл:

```text
internal/db_models/color.go
```

```go
package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Color struct {
	bun.BaseModel `bun:"table:colors"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull"`
}
```

в файл cmd/schema/main.go

добавляем &db_models.Color{}, по аналогии

```go
func main() {
	stmts, err := bunschema.New(
		bunschema.DialectPostgres,
	).Load(
		&db_models.Car{},
		&db_models.Color{},
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(stmts)
}

```

---

# 4. Сгенерировать актуальную схему

После создания модели:

```bash
go run ./cmd/schema > db/schema/schema.sql
```

Проверить изменения в:

```text
db/schema/schema.sql
```

Убедиться, что появилась таблица:

```sql
CREATE TABLE colors
```

---

# 5. Создать миграцию

Создать миграцию командой:

```bash
./cmd/schema/atlas.sh migrate diff add_colors_dictionary --env local
```

Проверить созданную миграцию.

Atlas должен создать миграцию с созданием таблицы:

```sql
CREATE TABLE colors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code TEXT NOT NULL
);
```

---

После генерации доработать миграцию вручную.

Добавить уникальный индекс:

```sql
CREATE UNIQUE INDEX idx_colors_code
ON colors(code);
```

Почему нужен индекс:

- поиск цвета по `code` будет выполняться перед каждым созданием
- гарантирует уникальность данных на уровне БД
- защищает от гонок запросов
- стоимость поддержки индекса для справочника практически нулевая

---

Добавить начальные данные:

```sql
INSERT INTO colors (code)
VALUES
('black'),
('white'),
('silver'),
('gray'),
('red'),
('blue'),
('green'),
('yellow'),
('orange'),
('brown'),
('beige'),
('gold'),
('purple');
```

Почему данные добавляются в миграцию:

- справочник должен быть готов сразу после применения миграции
- фронт может начать использовать цвета без ручного заполнения базы
- окружения будут создаваться одинаковыми

после изменения миграции вручную, нужно пересчитать контрольные суммы:

./cmd/schema/atlas.sh migrate hash --env local

---

Проверить итоговую миграцию перед применением.

# 6. Применить миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate apply --env local
```

Проверить результат.

Через pgAdmin:

```sql
SELECT *
FROM colors
ORDER BY code;
```

Либо через psql:

```sql
SELECT *
FROM colors
ORDER BY code;
```

Убедиться, что:

- таблица `colors` создана
- индекс `idx_colors_code` создан
- начальные данные вставлены

Дополнительно проверить:

```sql
SELECT COUNT(*)
FROM colors;
```

\ожидаем увидеть:

```text
13
```

если записей либо больше если список цветов был расширен.

- ***

# 7. Репозиторий

Создать:

```text
internal/repositories/color_repository.go
```

Интерфейс:

```go
type ColorRepository interface {
	GetAllRepository(ctx context.Context) ([]db_models.Color, error)

	GetByIDRepository(
		ctx context.Context,
		id uuid.UUID,
	) (*db_models.Color, error)

	GetByCodeRepository(
		ctx context.Context,
		code string,
	) (*db_models.Color, error)

	CreateRepository(
		ctx context.Context,
		color *db_models.Color,
	) error

	UpdateRepository(
		ctx context.Context,
		id uuid.UUID,
		request api_models.UpdateColorRequest,
	) (*db_models.Color, error)
}
```

---

# 8. Особенность POST

Перед созданием обязательно выполнить поиск по коду.

Псевдокод:

```go
existing, err := repo.GetByCodeRepository(
	ctx,
	request.Code,
)

if err == nil {
	return existing
}
```

Если запись найдена:

```text
не ошибка
не INSERT
возвращаем найденную запись
```

---

Далее обычный INSERT только если код отсутствует.

---

# 9. Валидация

Создать:

```text
internal/validation/color.go
```

Правила:

```go
ColorRules.Code.MinLength = 2
ColorRules.Code.MaxLength = 50
```

---

Create:

```go
ValidateCreateColorRequest()
```

Проверки:

```text
required
min_length
max_length
```

---

Update:

```go
ValidateUpdateColorRequest()
```

Проверять только присланные поля.

---

# 10. Сервисы

Создать:

```go
GetColorsService()
GetColorByIDService()
CreateColorService()
UpdateColorService()
```

---

# 11. Хендлеры

Создать:

```go
getColorsHandler()
getColorByIDHandler()
createColorHandler()
updateColorHandler()
```

---

# 12. Роутинг

Создать:

```go
ColorsMainHandler()
ColorByIDMainHandler()
```

По аналогии с Cars.

---

# 13. Проверочный набор curl

## Получить список

```bash
curl http://localhost:8080/colors
```

---

## Создать цвет

```bash
curl -X POST http://localhost:8080/colors \
-H "Content-Type: application/json" \
-d '{
  "code": "pink"
}'
```

---

## Повторное создание

```bash
curl -X POST http://localhost:8080/colors \
-H "Content-Type: application/json" \
-d '{
  "code": "pink"
}'
```

Ожидаем:

```text
без ошибки
тот же объект
новая запись не создаётся
```

---

## Получить по id

```bash
curl http://localhost:8080/colors/<id>
```

---

## Обновить

```bash
curl -X PATCH http://localhost:8080/colors/<id> \
-H "Content-Type: application/json" \
-d '{
  "code": "dark_blue"
}'
```

---

## Ошибка валидации

```bash
curl -X PATCH http://localhost:8080/colors/<id> \
-H "Content-Type: application/json" \
-d '{
  "code": "A"
}'
```

Ожидаем:

```json
{
  "errors": {
    "code": [
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

Работа считается завершённой, если:

- OpenAPI обновлён
- Go модели сгенерированы
- таблица colors создана
- индекс на code создан
- стартовые данные загружены
- GET список работает
- GET по id работает
- POST создаёт запись
- POST повторно возвращает существующую запись
- PATCH работает
- валидация работает
- набор curl проходит успешно
  ```\*\*````
