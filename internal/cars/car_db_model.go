package car

import (
	"time"

	car_body_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/body_type"
	car_brand "github.com/Adisey/car-rental-api/internal/cars/dictionaries/brand"
	car_color "github.com/Adisey/car-rental-api/internal/cars/dictionaries/color"
	car_drive_types "github.com/Adisey/car-rental-api/internal/cars/dictionaries/drive_type"
	car_fuel "github.com/Adisey/car-rental-api/internal/cars/dictionaries/fuel_type"
	car_model "github.com/Adisey/car-rental-api/internal/cars/dictionaries/model"
	car_transmission_type "github.com/Adisey/car-rental-api/internal/cars/dictionaries/transmission_type"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CarDBModel struct {
	bun.BaseModel `bun:"table:cars"`

	ID               uuid.UUID                                      `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name             string                                         `bun:"type:text,notnull"`
	Description      *string                                        `bun:"type:text"`
	BodyTypeID       *uuid.UUID                                     `bun:"type:uuid"`
	BodyType         *car_body_type.BodyTypeDBModel                 `bun:"rel:belongs-to,join:body_type_id=id"`
	PowerHp          *string                                        `bun:"type:integer"`
	BrandID          *uuid.UUID                                     `bun:"type:uuid"`
	Brand            *car_brand.BrandDBModel                        `bun:"rel:belongs-to,join:brand_id=id"`
	ColorID          *uuid.UUID                                     `bun:"type:uuid"`
	Color            *car_color.ColorDBModel                        `bun:"rel:belongs-to,join:color_id=id"`
	DriveID          *uuid.UUID                                     `bun:"type:uuid"`
	DriveType        *car_drive_types.DriveTypeDBModel              `bun:"rel:belongs-to,join:drive_id=id"`
	FuelID           *uuid.UUID                                     `bun:"type:uuid"`
	FuelType         *car_fuel.FuelTypeDBModel                      `bun:"rel:belongs-to,join:fuel_id=id"`
	ModelID          *uuid.UUID                                     `bun:"type:uuid"`
	Model            *car_model.ModelDBModel                        `bun:"rel:belongs-to,join:model_id=id"`
	TransmissionID   *uuid.UUID                                     `bun:"type:uuid"`
	TransmissionType *car_transmission_type.TransmissionTypeDBModel `bun:"rel:belongs-to,join:transmission_id=id"`
	CreatedAt        time.Time                                      `bun:",notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time                                      `bun:",notnull,default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time                                     `bun:"type:timestamptz"`
}
