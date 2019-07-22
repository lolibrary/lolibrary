package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadAttribute(ctx context.Context, id string) (*domain.Attribute, error) {
	snap, err := attributesByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("attribute", "id", id)
		}

		return nil, terrors.Wrap(err, nil)
	}

	attribute := &domain.Attribute{}
	if err := snap.DataTo(&attribute); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return attribute, nil
}

func ReadAttributeBySlug(ctx context.Context, slug string) (*domain.Attribute, error) {
	snap, err := attributesBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("attribute", "slug", slug)
		}

		return nil, terrors.Wrap(err, nil)
	}

	attribute := &domain.Attribute{}
	if err := snap.DataTo(&attribute); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return attribute, nil
}

func CreateAttribute(ctx context.Context, attribute *domain.Attribute) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(attributesByID.Doc(attribute.ID), attribute); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("attribute", "id", attribute.ID)
			}

			return err
		}

		if err := tx.Create(attributesBySlug.Doc(attribute.Slug), attribute); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("attribute", "slug", attribute.Slug)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateAttribute(ctx context.Context, attribute *domain.Attribute) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(attributesByID.Doc(attribute.ID), attribute); err != nil {
			return err
		}

		if err := tx.Set(attributesBySlug.Doc(attribute.Slug), attribute); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteAttribute(ctx context.Context, attribute *domain.Attribute) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(attributesByID.Doc(attribute.ID)); err != nil {
			return err
		}

		if err := tx.Delete(attributesBySlug.Doc(attribute.Slug)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListAttributes(ctx context.Context) ([]*domain.Attribute, error) {
	documents := attributesByID.Documents(ctx)
	defer documents.Stop()

	attributes := make([]*domain.Attribute, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		attribute := &domain.Attribute{}
		if err := doc.DataTo(&attribute); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		attributes = append(attributes, attribute)
	}

	return attributes, nil
}
