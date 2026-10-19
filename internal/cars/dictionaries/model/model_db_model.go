package car_model

import (
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ModelDBModel struct {
	bun.BaseModel `bun:"table:models"`
	ID            uuid.UUID               `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name          string                  `bun:"type:text,notnull,unique:idx_models_name_brand"`
	BrandID       *uuid.UUID              `bun:"type:uuid,unique:idx_models_name_brand"`
	Brand         *car_brand.BrandDBModel `bun:"rel:belongs-to,join:brand_id=id"`
}
