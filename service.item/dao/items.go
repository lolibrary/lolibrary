package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/monzo/terrors"
)

func ReadItem(id string) (*domain.Item, error) {
	i := &item{}

	res := DB.Where("id = ?", id).First(i)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return daoToDomain(i)
}

func ReadItemBySlug(slug string) (*domain.Item, error) {
	i := &item{}

	res := DB.Where("slug = ?", slug).First(i)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return daoToDomain(i)
}

func UpdateItem(input *domain.Item) error {
	i, err := domainToDAO(input)
	if err != nil {
		return err
	}

	res := DB.Table("items").Update(i)
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
