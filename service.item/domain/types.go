package domain

import (
	"time"
)

type Item struct {
	ID            string            `firestore:"id"`
	Slug          string            `firestore:"slug"`
	BrandID       string            `firestore:"brand_id"`
	CategoryID    string            `firestore:"category_id"`
	UserID        string            `firestore:"user_id"`
	ImageID       string            `firestore:"image_id"`
	PublisherID   string            `firestore:"publisher_id"`
	EnglishName   string            `firestore:"english_name"`
	ForeignName   string            `firestore:"foreign_name"`
	ProductNumber string            `firestore:"product_number"`
	Currency      string            `firestore:"currency"`
	Price         string            `firestore:"price"`
	Year          int               `firestore:"year"`
	Notes         string            `firestore:"notes"`
	Status        int               `firestore:"status"`
	Metadata      map[string]string `firestore:"metadata"`
	CreatedAt     time.Time         `firestore:"created_at"`
	UpdatedAt     time.Time         `firestore:"updated_at"`
	PublishedAt   time.Time         `firestore:"published_at"`
}
