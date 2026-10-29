package customers

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CustomerDBModel struct {
	bun.BaseModel `bun:"table:customers,alias:customer"`

	ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`

	FirstName string  `bun:"type:text,notnull"`
	LastName  *string `bun:"type:text"`

	Phone *string `bun:"type:text,"`
	Email *string `bun:"type:text,"`

	CreatedAt time.Time `bun:"created_at,notnull"`
	UpdatedAt time.Time `bun:"updated_at,notnull"`
}
