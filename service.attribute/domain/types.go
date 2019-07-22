package domain

import (
	"strings"
	"time"
)

type Attribute struct {
	// ID is a UUID.
	ID string `firestore:"id"`

	// Slug is a unique human-readable identifier used in URLs.
	Slug string `firestore:"slug"`

	// Name is the human readable name of this attribute.
	Name string `firestore:"name"`

	// CreatedAt is when this record was inserted.
	CreatedAt time.Time `firestore:"created_at"`

	// UpdatedAt is when this record was last updated.
	UpdatedAt time.Time `firestore:"updated_at"`
}

type AttributeValue struct {
	// ID is actually a compound of both attribute and item ID, to simulate unique(attribute_id, item_id).
	ID          string `firestore:"id"`
	ItemID      string `firestore:"item_id"`
	AttributeID string `firestore:"attribute_id"`
	Value       string `firestore:"value"`

	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}

// DocumentID will return a unique ID for this document.
func (av *AttributeValue) DocumentID() string {
	if av.ID == "" {
		av.ID = strings.Join([]string{av.ItemID, av.AttributeID}, "§")
	}

	return av.ID
}
