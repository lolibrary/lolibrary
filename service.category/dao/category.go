package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.category/domain"
	"github.com/monzo/terrors"
)

func CreateCategory(category *domain.Category) error {
	res := DB.Create(category)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateCategory(category *domain.Category) error {
	res := DB.Update(category)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("category", "Category not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadCategory(id string) (*domain.Category, error) {
	category := &domain.Category{}

	res := DB.Where("id = ?", id).First(category)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return category, nil
}

func ReadCategoryBySlug(slug string) (*domain.Category, error) {
	category := &domain.Category{}

	res := DB.Where("slug = ?", slug).First(category)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return category, nil
}

func DeleteCategory(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Category{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListCategories() ([]*domain.Category, error) {
	categories := make([]*domain.Category, 0)

	res := DB.Find(&categories)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return categories, nil
}
