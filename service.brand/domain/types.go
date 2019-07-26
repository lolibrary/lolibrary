package domain

import (
	"time"
)

// Brand is a representation of a fashion brand in the database.
type Brand struct {
	ID          string    `firestore:"id"`
	ImageID     string    `firestore:"image_id"`
	Slug        string    `firestore:"slug"`
	Name        string    `firestore:"name"`
	ShortName   string    `firestore:"short_name"`
	Description string    `firestore:"description"`
	CreatedAt   time.Time `firestore:"created_at"`
	UpdatedAt   time.Time `firestore:"updated_at"`
}
