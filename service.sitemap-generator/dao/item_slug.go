package dao

import (
	"context"

	"github.com/lolibrary/lolibrary/service.sitemap-generator/domain"
	"github.com/monzo/terrors"
)

func SitemapItemCursor(ctx context.Context, f func(item *domain.SitemapItem) error) (err error) {
	// todo: the .Where in here should be .Where("status = ?", itemproto.Status_PUBLISHED)

	rows, err := DB.Table("items").Where("status = 1").Select("slug, updated_at").Order("created_at asc").Rows()
	if err != nil {
		return terrors.Wrap(err, nil)
	}

	defer func() {
		rowErr := rows.Close()
		if err == nil {
			err = terrors.Wrap(rowErr, nil)
		}
	}()

	for rows.Next() {
		item := &domain.SitemapItem{}
		if err := rows.Scan(&item); err != nil {
			return terrors.Wrap(err, nil)
		}

		if err := f(item); err != nil {
			return terrors.Wrap(err, nil)
		}
	}

	return rows.Err()
}