package image

import (
	"os"
	"time"

	"github.com/monzo/typhon"
)

// cdn is a prefix to give image URLs.
// This is set by the IMAGE_CDN_URL environment variable once at boot.
var (
	cdn  = "https://minio.lolibrary.test/images/"
	opts = "?width=300&height=300&fit=bounds"
)

// Image is a database representation of an image in the CDN.
type Image struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Filename  string     `json:"filename" db:"filename"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

// URL returns a full URL to an image hosted in a CDN.
func (i Image) URL() string {
	return cdn + i.Filename
}

// Thumbnail returns a full URL to an image, with options to return a thumbnail.
func (i Image) Thumbnail() string {
	return cdn + i.Filename + opts
}

// NewService returns a new typhon Service.
func NewService() typhon.Service {
	c := os.Getenv("IMAGE_CDN_URL")
	if c != "" {
		cdn = c
	}

	// TODO: create s3 client/uploader

	return func(req typhon.Request) typhon.Response {
		return req.Response(nil)
	}
}
