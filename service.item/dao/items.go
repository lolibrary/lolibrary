package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/monzo/terrors"
)

func ReadItem(ctx context.Context, id string) (*domain.Item, error) {
	snap, err := itemsByID.Doc(id).Get(ctx)
	if err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	item := &domain.Item{}
	if err := snap.DataTo(&item); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return item, nil
}

func ReadItemBySlug(ctx context.Context, slug string) (*domain.Item, error) {
	snap, err := itemsBySlug.Doc(slug).Get(ctx)
	if err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	item := &domain.Item{}
	if err := snap.DataTo(&item); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return item, nil
}

func CreateItem(ctx context.Context, item *domain.Item) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(itemsByID.Doc(item.ID), item); err != nil {
			return err
		}

		if err := tx.Create(itemsBySlug.Doc(item.Slug), item); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}
