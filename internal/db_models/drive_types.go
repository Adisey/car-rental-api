package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DriveType struct {
	bun.BaseModel `bun:"table:drive_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_drive_type_code"`
}
