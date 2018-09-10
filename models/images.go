package models

import "time"

// Image is a database representation of an image in the CDN.
type Image struct {
	Model

	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Filename  string     `json:"filename"`
	Thumbnail string     `json:"thumbnail"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
}
