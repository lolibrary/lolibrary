package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.category/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadCategory(ctx context.Context, id string) (*domain.Category, error) {
	snap, err := categoriesByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("category", "id", id)
		}

		return nil, terrors.Wrap(err, nil)
	}

	category := &domain.Category{}
	if err := snap.DataTo(&category); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return category, nil
}

func ReadCategoryBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	snap, err := categoriesBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, database.ErrNotFound("category", "slug", slug)
		}

		return nil, terrors.Wrap(err, nil)
	}

	category := &domain.Category{}
	if err := snap.DataTo(&category); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return category, nil
}

func CreateCategory(ctx context.Context, category *domain.Category) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(categoriesByID.Doc(category.ID), category); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("category", "id", category.ID)
			}

			return err
		}

		if err := tx.Create(categoriesBySlug.Doc(category.Slug), category); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("category", "slug", category.Slug)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateCategory(ctx context.Context, category *domain.Category) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(categoriesByID.Doc(category.ID), category); err != nil {
			return err
		}

		if err := tx.Set(categoriesBySlug.Doc(category.Slug), category); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteCategory(ctx context.Context, category *domain.Category) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(categoriesByID.Doc(category.ID)); err != nil {
			return err
		}

		if err := tx.Delete(categoriesBySlug.Doc(category.Slug)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListCategories(ctx context.Context) ([]*domain.Category, error) {
	documents := categoriesByID.Documents(ctx)
	defer documents.Stop()

	categories := make([]*domain.Category, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		category := &domain.Category{}
		if err := doc.DataTo(&category); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		categories = append(categories, category)
	}

	return categories, nil
}
