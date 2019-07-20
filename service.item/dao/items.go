package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/monzo/terrors"
)

func ReadItem(id string) (*domain.Item, error) {
	item := &daoItem{}

	res := DB.Where("id = ?", id).First(item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return daoToDomain(item)
}

func ReadItemBySlug(slug string) (*domain.Item, error) {
	item := &daoItem{}

	res := DB.Where("slug = ?", slug).First(item)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return daoToDomain(item)
}

func UpdateItem(input *domain.Item) error {
	item, err := domainToDAO(input)
	if err != nil {
		return err
	}

	res := DB.Update(item)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("item", "Item not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}
