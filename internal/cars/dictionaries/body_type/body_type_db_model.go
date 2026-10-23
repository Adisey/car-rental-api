package car_body

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BodyTypeDBModel struct {
	bun.BaseModel `bun:"table:body_types,alias:body_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_body_type_code"`
}
