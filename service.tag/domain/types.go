package domain

import "time"

type Tag struct {
	// ID is a UUID.
	ID string

	// Slug is a unique human-readable identifier used in URLs.
	Slug string

	// Name is the human readable name of this tag.
	Name string

	// CreatedAt is when this record was inserted.
	CreatedAt time.Time

	// UpdatedAt is when this record was last updated.
	UpdatedAt time.Time
}
