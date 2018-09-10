package models

import "time"

// Attribute is a generic attribute on a model.
// The pivot between items and attributes contains the value of the attribute.
type Attribute struct {
	Model

	ID        string     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	Value     string     `json:"value"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// Category is the "type" of item, e.g. JSK, OP, Skirt.
type Category struct {
	Model

	ID        string     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// Feature is a representation of a feature on an item, e.g. back shirring, long sleeves.
type Feature struct {
	Model

	ID        string     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// Color is a colorway of an item, e.g. Black, White, Gold.
type Color struct {
	Model

	ID        string     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}

// Tag is a generic tag on an item used for searching.
type Tag struct {
	Model

	ID        string     `json:"id"`
	Slug      string     `json:"slug"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}
