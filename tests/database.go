package tests

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func PrepareDatabase() error {
	fmt.Println("=== TEST DATABASE SETUP ===")
	if err := CleanupDatabase(); err != nil {
		return err
	}
	cmd := exec.Command(
		"./cmd/schema/atlas.sh",
		"migrate",
		"apply",
		"--env",
		"test",
	)
	// Запускаем из корня проекта
	cmd.Dir = ".."
	// Добавляем переменную для atlas.sh
	cmd.Env = append(
		os.Environ(),
		"ENV_FILE=.env.test",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CleanupDatabase() error {
	fmt.Println("=== TEST DATABASE CLEANUP ===")
	println("CleanupDatabase - test TEST_DB_NAME")

	dbName := os.Getenv("TEST_DB_NAME")
	if !strings.Contains(dbName, "test") {
		return fmt.Errorf(
			"refusing to cleanup non-test database: %s",
			dbName,
		)
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("TEST_DB_HOST"),
		os.Getenv("TEST_DB_PORT"),
		os.Getenv("TEST_DB_NAME"),
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Printf(
		"Cleaning test database %s\n",
		os.Getenv("TEST_DB_NAME"),
	)
	_, err = db.Exec(`DROP SCHEMA IF EXISTS atlas_schema_revisions CASCADE;`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		DROP SCHEMA IF EXISTS public CASCADE;
		CREATE SCHEMA public;
	`)

	return err
}
