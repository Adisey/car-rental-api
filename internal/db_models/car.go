package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Car struct {
	bun.BaseModel `bun:"table:cars"`

	ID          uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name        string    `bun:"type:text,notnull"`
	Description *string   `bun:"type:text"`
}
