package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.color/domain"
	"github.com/monzo/terrors"
)

func CreateColor(color *domain.Color) error {
	res := DB.Create(color)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateColor(color *domain.Color) error {
	res := DB.Update(color)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("color", "Color not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadColor(id string) (*domain.Color, error) {
	color := &domain.Color{}

	res := DB.Where("id = ?", id).First(color)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return color, nil
}

func ReadColorBySlug(slug string) (*domain.Color, error) {
	color := &domain.Color{}

	res := DB.Where("slug = ?", slug).First(color)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return color, nil
}

func DeleteColor(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Color{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListColors() ([]*domain.Color, error) {
	colors := make([]*domain.Color, 0)

	res := DB.Find(&colors)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return colors, nil
}
