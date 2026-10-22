package tests

import (
	"fmt"
	"os"
	"os/exec"
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
		"local",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func CleanupDatabase() error {

	fmt.Println("=== TEST DATABASE CLEANUP ===")

	sql := `
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
`

	cmd := exec.Command(
		"psql",
		"-h", os.Getenv("DB_HOST"),
		"-p", os.Getenv("DB_PORT"),
		"-U", os.Getenv("DB_USER"),
		"-d", os.Getenv("DB_NAME"),
		"-c", sql,
	)

	cmd.Env = append(
		os.Environ(),
		"PGPASSWORD="+os.Getenv("DB_PASSWORD"),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
