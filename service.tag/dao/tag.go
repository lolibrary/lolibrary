package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.tag/domain"
	"github.com/monzo/terrors"
)

func CreateTag(tag *domain.Tag) error {
	res := DB.Create(tag)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateTag(tag *domain.Tag) error {
	res := DB.Update(tag)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("tag", "Tag not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadTag(id string) (*domain.Tag, error) {
	tag := &domain.Tag{}

	res := DB.Where("id = ?", id).First(tag)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return tag, nil
}

func ReadTagBySlug(slug string) (*domain.Tag, error) {
	tag := &domain.Tag{}

	res := DB.Where("slug = ?", slug).First(tag)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return tag, nil
}

func DeleteTag(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Tag{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListTags() ([]*domain.Tag, error) {
	tags := make([]*domain.Tag, 0)

	res := DB.Find(&tags)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return tags, nil
}
