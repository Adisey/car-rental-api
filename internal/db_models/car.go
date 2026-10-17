package db_models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Car struct {
	bun.BaseModel `bun:"table:cars"`

	ID               uuid.UUID         `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name             string            `bun:"type:text,notnull"`
	Description      *string           `bun:"type:text"`
	BodyTypeID       *uuid.UUID        `bun:"type:uuid"`
	BodyType         *BodyType         `bun:"rel:belongs-to,join:body_type_id=id"`
	PowerHp          *string           `bun:"type:integer"`
	BrandID          *uuid.UUID        `bun:"type:uuid"`
	Brand            *Brand            `bun:"rel:belongs-to,join:brand_id=id"`
	ColorID          *uuid.UUID        `bun:"type:uuid"`
	Color            *Color            `bun:"rel:belongs-to,join:color_id=id"`
	DriveID          *uuid.UUID        `bun:"type:uuid"`
	DriveType        *DriveType        `bun:"rel:belongs-to,join:drive_id=id"`
	FuelID           *uuid.UUID        `bun:"type:uuid"`
	FuelType         *FuelType         `bun:"rel:belongs-to,join:fuel_id=id"`
	ModelID          *uuid.UUID        `bun:"type:uuid"`
	Model            *Model            `bun:"rel:belongs-to,join:model_id=id"`
	TransmissionID   *uuid.UUID        `bun:"type:uuid"`
	TransmissionType *TransmissionType `bun:"rel:belongs-to,join:transmission_id=id"`
	CreatedAt        time.Time         `bun:",notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time         `bun:",notnull,default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time        `bun:"type:timestamptz"`
}
