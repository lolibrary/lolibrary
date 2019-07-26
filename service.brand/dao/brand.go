package dao

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ReadBrand(ctx context.Context, id string) (*domain.Brand, error) {
	snap, err := brandsByID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, nil
		}

		return nil, terrors.Wrap(err, nil)
	}

	brand := &domain.Brand{}
	if err := snap.DataTo(&brand); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return brand, nil
}

func ReadBrandBySlug(ctx context.Context, slug string) (*domain.Brand, error) {
	snap, err := brandsBySlug.Doc(slug).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, nil
		}

		return nil, terrors.Wrap(err, nil)
	}

	brand := &domain.Brand{}
	if err := snap.DataTo(&brand); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return brand, nil
}

func ReadBrandByShortName(ctx context.Context, name string) (*domain.Brand, error) {
	snap, err := brandsByShortName.Doc(name).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, nil
		}

		return nil, terrors.Wrap(err, nil)
	}

	brand := &domain.Brand{}
	if err := snap.DataTo(&brand); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return brand, nil
}

func CreateBrand(ctx context.Context, brand *domain.Brand) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Create(brandsByID.Doc(brand.ID), brand); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("brand", "id", brand.ID)
			}

			return err
		}

		if err := tx.Create(brandsBySlug.Doc(brand.Slug), brand); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("brand", "slug", brand.Slug)
			}

			return err
		}

		if err := tx.Create(brandsByShortName.Doc(brand.ShortName), brand); err != nil {
			if database.AlreadyExists(err) {
				return database.ErrAlreadyExists("brand", "short_name", brand.ShortName)
			}

			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateBrand(ctx context.Context, brand *domain.Brand) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Set(brandsByID.Doc(brand.ID), brand); err != nil {
			return err
		}

		if err := tx.Set(brandsBySlug.Doc(brand.Slug), brand); err != nil {
			return err
		}

		if err := tx.Set(brandsByShortName.Doc(brand.ShortName), brand); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteBrand(ctx context.Context, brand *domain.Brand) error {
	if err := Firestore.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := tx.Delete(brandsByID.Doc(brand.ID)); err != nil {
			return err
		}

		if err := tx.Delete(brandsBySlug.Doc(brand.Slug)); err != nil {
			return err
		}

		if err := tx.Delete(brandsByShortName.Doc(brand.ShortName)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func ListBrands(ctx context.Context) ([]*domain.Brand, error) {
	documents := brandsByID.Documents(ctx)
	defer documents.Stop()

	brands := make([]*domain.Brand, 0, 32)
	for {
		doc, err := documents.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		brand := &domain.Brand{}
		if err := doc.DataTo(&brand); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		brands = append(brands, brand)
	}

	return brands, nil
}
