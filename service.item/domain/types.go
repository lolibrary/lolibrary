package domain

import "time"

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

	Status   int
	Metadata map[string]string

	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
}

/*
string id = 1;
    string slug = 2;
    string brand_id = 3;
    string category_id = 4;
    string user_id = 5;
    string image_id = 6;
    string publisher_id = 7;
    string english_name = 8;
    string foreign_name = 9;
    string product_number = 10;
    string currency = 11;
    string price = 12;
    int32 year = 13;
    string notes = 14;

    // status is used for draft/non-draft but could be a bitmask later.
    int32 status = 40;

    // key-value searchable metadata.
    map<string, string> metadata = 50;

    // timestamps
    string created_at = 100;
    string updated_at = 101;
    string published_at = 102;
*/
