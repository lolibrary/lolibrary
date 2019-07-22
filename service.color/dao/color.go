package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.color/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadColor(ctx context.Context, id string) (*domain.Color, error) {
	snap, err := colorsByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("color", "id", id)
		}

		return nil, terrors.Wrap(err, nil)
	}

	color := &domain.Color{}
	if err := snap.DataTo(&color); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return color, nil
}

func ReadColorBySlug(ctx context.Context, slug string) (*domain.Color, error) {
	snap, err := colorsBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("color", "slug", slug)
		}

		return nil, terrors.Wrap(err, nil)
	}

	color := &domain.Color{}
	if err := snap.DataTo(&color); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return color, nil
}

func CreateColor(ctx context.Context, color *domain.Color) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(colorsByID.Doc(color.ID), color); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("color", "id", color.ID)
			}

			return err
		}

		if err := tx.Create(colorsBySlug.Doc(color.Slug), color); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("color", "slug", color.Slug)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateColor(ctx context.Context, color *domain.Color) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(colorsByID.Doc(color.ID), color); err != nil {
			return err
		}

		if err := tx.Set(colorsBySlug.Doc(color.Slug), color); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteColor(ctx context.Context, color *domain.Color) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(colorsByID.Doc(color.ID)); err != nil {
			return err
		}

		if err := tx.Delete(colorsBySlug.Doc(color.Slug)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListColors(ctx context.Context) ([]*domain.Color, error) {
	documents := colorsByID.Documents(ctx)
	defer documents.Stop()

	colors := make([]*domain.Color, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		color := &domain.Color{}
		if err := doc.DataTo(&color); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		colors = append(colors, color)
	}

	return colors, nil
}
