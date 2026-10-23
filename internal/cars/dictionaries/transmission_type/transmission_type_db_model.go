package car_transmission

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TransmissionTypeDBModel struct {
	bun.BaseModel `bun:"table:transmission_types,alias:transmission_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_transmission_type_code"`
}
