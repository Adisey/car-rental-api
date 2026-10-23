package car_drive_types

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DriveTypeDBModel struct {
	bun.BaseModel `bun:"table:drive_types,alias:drive_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_drive_type_code"`
}
