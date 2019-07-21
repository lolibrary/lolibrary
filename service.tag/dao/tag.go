package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.tag/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadTag(ctx context.Context, id string) (*domain.Tag, error) {
	snap, err := tagsByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("tag", "id", id)
		}

		return nil, terrors.Wrap(err, nil)
	}

	tag := &domain.Tag{}
	if err := snap.DataTo(&tag); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return tag, nil
}

func ReadTagBySlug(ctx context.Context, slug string) (*domain.Tag, error) {
	snap, err := tagsBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("tag", "slug", slug)
		}

		return nil, terrors.Wrap(err, nil)
	}

	tag := &domain.Tag{}
	if err := snap.DataTo(&tag); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return tag, nil
}

func CreateTag(ctx context.Context, tag *domain.Tag) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(tagsByID.Doc(tag.ID), tag); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("tag", "id", tag.ID)
			}

			return err
		}

		if err := tx.Create(tagsBySlug.Doc(tag.Slug), tag); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("tag", "slug", tag.Slug)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateTag(ctx context.Context, tag *domain.Tag) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(tagsByID.Doc(tag.ID), tag); err != nil {
			return err
		}

		if err := tx.Set(tagsBySlug.Doc(tag.Slug), tag); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteTag(ctx context.Context, tag *domain.Tag) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(tagsByID.Doc(tag.ID)); err != nil {
			return err
		}

		if err := tx.Delete(tagsBySlug.Doc(tag.Slug)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListTags(ctx context.Context) ([]*domain.Tag, error) {
	documents := tagsByID.Documents(ctx)
	defer documents.Stop()

	tags := make([]*domain.Tag, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		tag := &domain.Tag{}
		if err := doc.DataTo(&tag); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
