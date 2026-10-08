package db_models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Car struct {
	bun.BaseModel `bun:"table:cars"`

	ID          uuid.UUID  `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name        string     `bun:"type:text,notnull"`
	Description *string    `bun:"type:text"`
	ColorID     *uuid.UUID `bun:"type:uuid"`
	Color       *Color     `bun:"rel:belongs-to,join:color_id=id"`
	CreatedAt   time.Time  `bun:",notnull,default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `bun:",notnull,default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `bun:"type:timestamptz"`
}
