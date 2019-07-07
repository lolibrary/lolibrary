package domain

import "time"

type Attribute struct {
	// ID is a UUID.
	ID string

	// Slug is a unique human-readable identifier used in URLs.
	Slug string

	// Name is the human readable name of this attribute.
	Name string

	// CreatedAt is when this record was inserted.
	CreatedAt time.Time

	// UpdatedAt is when this record was last updated.
	UpdatedAt time.Time
}

type AttributeLink struct {
	BrandID   string
	Value     string
	Attribute Attribute
}
