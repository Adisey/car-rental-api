## Создаём таблицы

### internal/customers/customer_db_model.go

```go
type CustomerDBModel struct {
	bun.BaseModel `bun:"table:customers,alias:customer"`
	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	FirstName string  `bun:"type:text,notnull"`
	LastName  *string `bun:"type:text"`
	Phone *string `bun:"type:text,"`
	Email *string `bun:"type:text,"`
	CreatedAt time.Time `bun:"created_at,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
}

```

### cmd/schema/main.go

добавляем

```go
		&customers.CustomerDBModel{},
```

### Сгенерировать актуальную схему

Запустить:

```bash
go run ./cmd/schema >*db/schema/schema.sql
```

проверяем в db/schema/schema.sql

```sql
CREATE TABLE "customers"
```

### Создать миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate diff add_customer --env local
```

проверяем

```text
db/migrations/XXXXXXXXXXXXX_add_customer.sql
```

#### Если осознанно менял миграции вручную либо нужно перегенерировать её , то нужно пересчитать контрольные суммы:

```bash
./cmd/schema/atlas.sh migrate hash --env local
```

### Применить миграцию

Запустить:

```bash
./cmd/schema/atlas.sh migrate apply --env local
```

проверяем фал миграции

## типизируем API

```text
openapi/api.yaml
```

добавляем

components:
и
paths:

### Сгенерировать Go API модели

Запустить:

```bash
./cmd/openapi/generate.sh
```

### Проверяем результат

```text
internal/api_models/types.gen.go
```

## Создать репозиторий

Файлы:

```text
internal/customers/customer_repository.go
internal/customers/customer_repository_bun.go
```

## Создать сервисы

Файл:

```text
internal/customers/customer_service.go
```

Создать:

```go
GetCustomersService()
GetCustomerByIDService()
CreateCustomerService()
UpdateCustomerService()
```

### Создать валидацию

Файл:

```text
internal/customers/customer_validation.go
```

Создать:

```go
ValidateCreateCustomerRequest()
ValidateUpdateCustomerRequest()
```

Проверки:

```text
required
min_length
max_length
```

## Создать Хендлеры

Файл:

```text
internal/customers/customer_handler.go
```

### подключить в майн Хендлер

```text
cmd/server/main.go
```

## Пишем тесты
