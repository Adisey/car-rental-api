package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FuelType struct {
	bun.BaseModel `bun:"table:fuel_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_fuel_type_code"`
}
