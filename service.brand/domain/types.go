package domain

import (
	"time"
)

// Brand is a representation of a fashion brand in the database.
type Brand struct {
	// ID is a UUID uniquely identifying this brand.
	ID string

	// ImageID is used to link to an image for this brand.
	// See service.image.
	ImageID string

	// Slug is a human-readable unique identifier.
	// Slug is used in URLs and is hyphen-separated and normalized before storing/querying.
	Slug        string
	Name        string
	ShortName   string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
