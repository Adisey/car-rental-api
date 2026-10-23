package car_fuel

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FuelTypeDBModel struct {
	bun.BaseModel `bun:"table:fuel_types,alias:fuel_types"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Code string    `bun:"type:text,notnull,unique:idx_fuel_type_code"`
}
