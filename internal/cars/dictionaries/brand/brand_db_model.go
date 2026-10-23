package car_brand

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BrandDBModel struct {
	bun.BaseModel `bun:"table:brands,alias:brands"`

	ID   uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:"type:text,notnull,unique:idx_brands_name"`
}
