package models

import "time"

type Item struct {
	Model

	ID            string       `json:"id"`
	EnglishName   string       `json:"english_name"`
	ForeignName   string       `json:"foreign_name"`
	Slug          string       `json:"slug"`
	Brand         *Brand       `json:"brand"`
	Category      *Category    `json:"category"`
	ProductNumber string       `json:"product_number"`
	Year          int          `json:"year"`
	Notes         string       `json:"notes"`
	Features      []*Feature   `json:"features"`
	Colors        []*Color     `json:"colors"`
	Tags          []*Tag       `json:"tags"`
	Attributes    []*Attribute `json:"attribute"`
	Status        ItemStatus   `json:"status"`
	Price         string       `json:"price"`
	Currency      string       `json:"currency"`
	UpdatedAt     *time.Time   `json:"updated_at"`
	CreatedAt     *time.Time   `json:"created_at"`
}

// Status is an enum representation of an item's status.
type Status int

// ItemStatus enum values are draft/published.
const (
	StatusDraft     = 0
	StatusPublished = 1
)

// ItemStatus is an external GraphQL representation of ItemStatus.
type ItemStatus string

// ItemStatus graphql enum values.
const (
	ItemStatusDraft     = "DRAFT"
	ItemStatusPublished = "PUBLISHED"
)

// ItemStatuses is a map of Status (int) to ItemStatus (graphql string).
var ItemStatuses = map[Status]ItemStatus{
	StatusDraft:     ItemStatusDraft,
	StatusPublished: ItemStatusPublished,
}

// Statuses is a map of ItemStatus (graphql string) to Status (int)
var Statuses = map[ItemStatus]Status{
	ItemStatusDraft:     StatusDraft,
	ItemStatusPublished: StatusPublished,
}
