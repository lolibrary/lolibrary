package domain

import (
	"time"
)

type Item struct {
	ID            string
	Slug          string
	BrandID       string
	CategoryID    string
	UserID        string
	ImageID       string
	PublisherID   string
	EnglishName   string
	ForeignName   string
	ProductNumber string
	Currency      string
	Price         string
	Year          int
	Notes         string
	Status        int
	Metadata      map[string]string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   time.Time
}
