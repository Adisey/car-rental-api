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
