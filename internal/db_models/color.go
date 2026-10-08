package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Color struct {
	bun.BaseModel `bun:"table:colors"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_colors_code"`
}
