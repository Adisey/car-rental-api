# Тестирование

В проекте используются два вида тестов:

1. Обычные тесты (unit-тесты и тесты без реальной БД).
2. Интеграционные тесты (с реальной PostgreSQL базой данных).

---

# Обычные тесты

Запуск:

```bash
go test -count=1 -v ./tests/...
```

Особенности:

- Не используют отдельную тестовую БД.
- Выполняются быстро.
- Не накатывают миграции.
- Не требуют предварительной подготовки базы данных.

---

# Интеграционные тесты

Запуск:

```bash
go test -count=1 -v -tags=integration ./tests/...
```

или только для интеграционного каталога:

```bash
go test -count=1 -v -tags=integration ./tests/integration/...
```

---

# Что происходит при запуске интеграционных тестов

Перед выполнением тестов автоматически выполняется следующий сценарий:

1. Загружается конфигурация из `.env.test`.
2. Выполняется очистка тестовой базы данных.
3. Накатываются все миграции через Atlas.
4. Создаётся подключение Bun к тестовой БД.
5. Запускаются интеграционные тесты.

Схема запуска:

```text
Load .env.test
        ↓
CleanupDatabase()
        ↓
Atlas migrate apply --env test
        ↓
ConnectBun()
        ↓
Run Tests
```

---

# Тестовая база данных

Интеграционные тесты используют отдельную базу данных:

```text
car_rental_test
```

Тестовая база полностью изолирована от базы разработки.

Перед каждым запуском тестов база очищается и создаётся заново посредством наката всех миграций.

После завершения тестов база НЕ очищается.

Это сделано специально, чтобы:

- анализировать состояние данных после выполнения тестов;
- упрощать отладку;
- вручную проверять результаты через DBeaver, DataGrip и другие инструменты.

---

# Важно

Если после изменения структуры БД интеграционные тесты начали падать:

1. Проверьте, что новая миграция создана.
2. Проверьте, что миграция успешно выполняется командой:

```bash
./cmd/schema/atlas.sh migrate apply --env test
```

3. Повторно запустите интеграционные тесты.

---

# Добавление новых интеграционных тестов

Рекомендуемый подход:

Для сервисного слоя:

```text
Create
GetAll
GetByID
Update
Негативные сценарии
```

Для HTTP-обработчиков:

```text
GET     happy path
POST    happy path
PATCH   happy path

GET     invalid id
POST    invalid body
POST    validation error
PATCH   invalid body
PATCH   validation error
```

---

# Отключение кеша Go

Во всех командах используется:

```bash
-count=1
```

Это необходимо для того, чтобы Go не использовал кеш результатов тестов.

Без этого флага может показаться, что тесты выполнились, хотя фактически был использован сохранённый результат предыдущего запуска:

```text
(cached)
```

Для интеграционных тестов использование кеша недопустимо.

# Запуск отдельных тестов

## Запустить один файл с тестами

Например:

```bash
go test -count=1 -v -tags=integration ./tests/integration/body_type_integration_test.go
```

Однако такой способ используется редко, потому что Go может не собрать зависимости внутри пакета.

Обычно лучше запускать пакет целиком:

```bash
go test -count=1 -v -tags=integration ./tests/integration
```

---

## Запустить один тест

Флаг `-run` принимает регулярное выражение.

Запуск конкретного теста:

```bash
go test -count=1 -v -tags=integration ./tests/integration \
    -run Test_BodyType_Create_GetAll_Update_GetByID
```

---

## Запустить несколько тестов

Например все тесты для BodyType:

```bash
go test -count=1 -v -tags=integration ./tests/integration \
    -run BodyType
```

Будут выполнены:

```text
Test_BodyType_Create_GetAll_Update_GetByID
Test_BodyType_Create_DuplicateCode
Test_BodyType_Update_DuplicateCode
Test_BodyType_GetByID_NotFound

Test_HTTP_BodyType_GetAll
Test_HTTP_BodyType_Create
...
```

---

## Запустить только HTTP-тесты

```bash
go test -count=1 -v -tags=integration ./tests/integration \
    -run HTTP
```

---

## Запустить только сервисные тесты

```bash
go test -count=1 -v -tags=integration ./tests/integration \
    -run '^Test_BodyType'
```

---

## Показать список тестов без запуска

```bash
go test -tags=integration ./tests/integration -list .
```

---

## Полезно при разработке

Самая частая команда для отладки одного теста:

```bash
go test -count=1 -v -tags=integration ./tests/integration \
    -run Test_BodyType_Create_GetAll_Update_GetByID
```

Она:

- пересоздаёт тестовую БД;
- накатывает миграции;
- запускает только один нужный тест;
- выводит подробный лог.
