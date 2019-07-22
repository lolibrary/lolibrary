package dao

import (
	"context"

	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/monzo/terrors"
	"google.golang.org/api/iterator"
)

func ListAttributesByItem(ctx context.Context, itemID string) ([]*domain.AttributeValue, error) {
	docs := attributesByItemID.Where("item_id", "==", itemID).Documents(ctx)
	defer docs.Stop()

	values := make([]*domain.AttributeValue, 0)
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		av := &domain.AttributeValue{}
		if err := doc.DataTo(&av); err != nil {
			return nil, terrors.Wrap(err, nil)
		}

		values = append(values, av)
	}

	return values, nil
}

func ReadAttributeValue(ctx context.Context, id string) (*domain.AttributeValue, error) {
	snap, err := attributesByItemID.Doc(id).Get(ctx)
	if err != nil {
		if database.NotFound(err) {
			return nil, nil
		}

		return nil, terrors.Wrap(err, nil)
	}

	val := &domain.AttributeValue{}
	if err := snap.DataTo(&val); err != nil {
		return nil, terrors.Wrap(err, nil)
	}

	return val, nil
}

func CreateAttributeValue(ctx context.Context, val *domain.AttributeValue) error {
	if _, err := attributesByItemID.Doc(val.DocumentID()).Create(ctx, val); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func UpdateAttributeValue(ctx context.Context, val *domain.AttributeValue) error {
	if _, err := attributesByItemID.Doc(val.DocumentID()).Set(ctx, val); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}

func DeleteAttributeValue(ctx context.Context, val *domain.AttributeValue) error {
	if _, err := attributesByItemID.Doc(val.DocumentID()).Delete(ctx); err != nil {
		return terrors.Wrap(err, nil)
	}

	return nil
}
