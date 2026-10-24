# Testing

## Purpose

The project uses Go build tags to separate:

- Unit tests
- Integration tests
- End-to-End (E2E) tests

This allows fast tests to run by default and more expensive tests to be executed only when requested.

---

# Directory Structure

```text
tests/
├─ testmain_test.go
├─ database.go
├─ integration_test.go
└─ e2e_test.go
```

---

# Files

## testmain_test.go

Common entry point for Integration and E2E tests.

Responsible for:

1. Database setup
2. Running tests
3. Database cleanup

Current implementation only prints messages.

Future implementation:

```text
Drop schema
↓
Apply migrations
↓
Run tests
↓
Drop schema
```

---

## database.go

Contains:

```go
PrepareDatabase()
CleanupDatabase()
```

Current implementation only prints messages.

Future implementation will:

- connect to test database
- recreate schema
- apply Atlas migrations
- cleanup schema after test execution

---

## integration_test.go

Example Integration test.

Build tag:

```go
//go:build integration
```

Runs only for:

```bash
go test -tags=integration ./...
```

Если мы хотим без кеша и с повторым созданием БД и миграцией

```bash
go test -count=1 -v -tags=integration ./tests/...
```

---

## e2e_test.go

Example End-to-End test.

Build tag:

```go
//go:build e2e
```

Runs only for:

```bash
go test -tags=e2e ./...
```

---

# Commands

## Run Unit Tests

```bash
go test ./...
```

Runs only tests without build tags.

---

## Run Integration Tests

```bash
go test -v -tags=integration ./...
```

Expected output:

```text
=== DATABASE SETUP ===

INTEGRATION TEST EXECUTED

=== DATABASE CLEANUP ===
```

---

## Run E2E Tests

```bash
go test -v -tags=e2e ./...
```

Expected output:

```text
=== DATABASE SETUP ===

E2E TEST EXECUTED

=== DATABASE CLEANUP ===
```

---

# Future Database Lifecycle

Integration and E2E tests will use the same lifecycle:

```text
PrepareDatabase()

↓
Drop schema

↓
Apply all migrations

↓
Run tests

↓
CleanupDatabase()

↓
Drop schema
```

This guarantees that every test run starts from a clean database state.
