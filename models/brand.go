package models

import "time"

// Brand is an item brand, attached to every item.
type Brand struct {
	Model

	ID          string     `json:"id"`
	ShortName   string     `json:"short_name"`
	Description string     `json:"description"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CreatedAt   *time.Time `json:"created_at"`
}
