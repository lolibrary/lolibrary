package domain

import "time"

type Image struct {
	ID         string    `firestore:"id"`
	ItemID     string    `firestore:"item_id"`
	UploaderID string    `firestore:"uploader_id"`
	Filename   string    `firestore:"filename"`
	Type       string    `firestore:"type"`
	Uploaded   bool      `firestore:"uploaded"`
	CreatedAt  time.Time `firestore:"created_at"`
	UpdatedAt  time.Time `firestore:"updated_at"`
}
