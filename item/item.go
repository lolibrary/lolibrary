package item

import (
	"fmt"
	"time"

	"github.com/lolibrary/lolibrary/currency"
)

type Item struct {
	ID            string            `json:"id"`
	EnglishName   string            `json:"english_name"`
	ForeignName   string            `json:"foreign_name"`
	Slug          string            `json:"slug"`
	BrandID       string            `json:"brand_id"`
	CategoryID    string            `json:"category_id"`
	ImageID       string            `json:"image_id"`
	ProductNumber string            `json:"product_number"`
	Year          int               `json:"year"`
	Notes         string            `json:"notes"`
	Status        Status            `json:"status"`
	Price         string            `json:"price"`
	Currency      currency.Currency `json:"currency"`
	UpdatedAt     *time.Time        `json:"updated_at"`
	CreatedAt     *time.Time        `json:"created_at"`
}

// Status is an enum representation of an Item's status.
type Status int

const (
	// StatusDraft indicates an item that is currently a draft.
	StatusDraft Status = 0

	// StatusPublished indicates an item that is published.
	StatusPublished Status = 1
)

// String returns a string representation of a Status.
func (s Status) String() string {
	switch s {
	case StatusPublished:
		return "published"
	case StatusDraft:
		return "draft"
	}

	return fmt.Sprintf("unknown status: %d", s)
}
