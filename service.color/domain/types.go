package domain

import "time"

type Color struct {
	// ID is a UUID.
	ID string `firestore:"id"`

	// Slug is a unique human-readable identifier used in URLs.
	Slug string `firestore:"slug"`

	// Name is the human readable name of this color.
	Name string `firestore:"name"`

	// CreatedAt is when this record was inserted.
	CreatedAt time.Time `firestore:"created_at"`

	// UpdatedAt is when this record was last updated.
	UpdatedAt time.Time `firestore:"updated_at"`
}
