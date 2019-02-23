package brand

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Brand is an item brand, attached to every item.
type Brand struct {
	ID          string     `json:"id"`
	Slug        string     `json:"slug"`
	ShortName   string     `json:"short_name"`
	Description string     `json:"description"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CreatedAt   *time.Time `json:"created_at"`
}

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) QueryByUUID(uuid string) *Brand {
	var brand Brand

	r.db.
}