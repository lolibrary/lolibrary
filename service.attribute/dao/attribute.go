package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	"github.com/monzo/terrors"
)

func CreateAttribute(attribute *domain.Attribute) error {
	res := DB.Create(attribute)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateAttribute(attribute *domain.Attribute) error {
	res := DB.Update(attribute)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("attribute", "Attribute not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadAttribute(id string) (*domain.Attribute, error) {
	attribute := &domain.Attribute{}

	res := DB.Where("id = ?", id).First(attribute)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return attribute, nil
}

func ReadAttributeBySlug(slug string) (*domain.Attribute, error) {
	attribute := &domain.Attribute{}

	res := DB.Where("slug = ?", slug).First(attribute)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return attribute, nil
}

func DeleteAttribute(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Attribute{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListAttributes() ([]*domain.Attribute, error) {
	attributes := make([]*domain.Attribute, 0)

	res := DB.Find(&attributes)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return attributes, nil
}
