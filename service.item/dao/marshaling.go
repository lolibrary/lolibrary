package dao

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/monzo/terrors"
)

type item struct {
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
	Metadata      postgres.Jsonb
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PublishedAt   time.Time
}

// TableName sets this model's table name to "users", overriding the default (dao_users).
func (item) TableName() string {
	return "items"
}

func domainToDAO(model *domain.Item) (*item, error) {
	// first we need to marshal our metadata to []byte
	metadata := []byte("{}")
	if model.Metadata != nil {
		b, err := json.Marshal(model.Metadata)
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}
		metadata = b
	}

	return &item{
		ID:            model.ID,
		Slug:          model.Slug,
		BrandID:       model.BrandID,
		CategoryID:    model.CategoryID,
		UserID:        model.UserID,
		ImageID:       model.ImageID,
		PublisherID:   model.PublisherID,
		EnglishName:   model.EnglishName,
		ForeignName:   model.ForeignName,
		ProductNumber: model.ProductNumber,
		Currency:      model.Currency,
		Price:         model.Price,
		Year:          model.Year,
		Notes:         model.Notes,
		Status:        model.Status,
		Metadata:      postgres.Jsonb{RawMessage: metadata},
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		PublishedAt:   model.PublishedAt,
	}, nil
}

func daoToDomain(model *item) (*domain.Item, error) {
	// decode the raw message
	var metadata map[string]string
	if err := json.Unmarshal(model.Metadata.RawMessage, &metadata); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return &domain.Item{
		ID:            model.ID,
		Slug:          model.Slug,
		BrandID:       model.BrandID,
		CategoryID:    model.CategoryID,
		UserID:        model.UserID,
		ImageID:       model.ImageID,
		PublisherID:   model.PublisherID,
		EnglishName:   model.EnglishName,
		ForeignName:   model.ForeignName,
		ProductNumber: model.ProductNumber,
		Currency:      model.Currency,
		Price:         model.Price,
		Year:          model.Year,
		Notes:         model.Notes,
		Status:        model.Status,
		Metadata:      metadata,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		PublishedAt:   model.PublishedAt,
	}, nil
}
