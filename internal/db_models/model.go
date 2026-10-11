package db_models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Model struct {
	bun.BaseModel `bun:"table:models"`
	ID            uuid.UUID  `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name          string     `bun:"type:text,notnull,unique:idx_models_name_brand"`
	BrandID       *uuid.UUID `bun:"type:uuid,unique:idx_models_name_brand"`
	Brand         *Brand     `bun:"rel:belongs-to,join:brand_id=id"`
}
