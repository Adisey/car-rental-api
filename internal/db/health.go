package db

import (
	"context"
	"time"
)

func CheckHealth(ctx context.Context) error {
	var currentTime time.Time

	err := Pool.QueryRow(
		ctx,
		"SELECT NOW()",
	).Scan(&currentTime)

	return err
}
