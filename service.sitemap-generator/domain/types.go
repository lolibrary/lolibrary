package domain

import "time"

type SitemapItem struct {
	Slug      string
	UpdatedAt time.Time
}
