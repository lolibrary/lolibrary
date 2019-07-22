package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadFeature(ctx context.Context, id string) (*domain.Feature, error) {
	snap, err := featuresByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("feature", "id", id)
		}

		return nil, terrors.Wrap(err, nil)
	}

	feature := &domain.Feature{}
	if err := snap.DataTo(&feature); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return feature, nil
}

func ReadFeatureBySlug(ctx context.Context, slug string) (*domain.Feature, error) {
	snap, err := featuresBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("feature", "slug", slug)
		}

		return nil, terrors.Wrap(err, nil)
	}

	feature := &domain.Feature{}
	if err := snap.DataTo(&feature); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return feature, nil
}

func CreateFeature(ctx context.Context, feature *domain.Feature) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(featuresByID.Doc(feature.ID), feature); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("feature", "id", feature.ID)
			}

			return err
		}

		if err := tx.Create(featuresBySlug.Doc(feature.Slug), feature); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("feature", "slug", feature.Slug)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateFeature(ctx context.Context, feature *domain.Feature) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(featuresByID.Doc(feature.ID), feature); err != nil {
			return err
		}

		if err := tx.Set(featuresBySlug.Doc(feature.Slug), feature); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteFeature(ctx context.Context, feature *domain.Feature) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(featuresByID.Doc(feature.ID)); err != nil {
			return err
		}

		if err := tx.Delete(featuresBySlug.Doc(feature.Slug)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListFeatures(ctx context.Context) ([]*domain.Feature, error) {
	documents := featuresByID.Documents(ctx)
	defer documents.Stop()

	features := make([]*domain.Feature, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		feature := &domain.Feature{}
		if err := doc.DataTo(&feature); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		features = append(features, feature)
	}

	return features, nil
}
