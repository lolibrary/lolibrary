package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	"github.com/monzo/terrors"
)

func CreateFeature(feature *domain.Feature) error {
	res := DB.Create(feature)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateFeature(feature *domain.Feature) error {
	res := DB.Update(feature)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("feature", "Feature not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadFeature(id string) (*domain.Feature, error) {
	feature := &domain.Feature{}

	res := DB.Where("id = ?", id).First(feature)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return feature, nil
}

func ReadFeatureBySlug(slug string) (*domain.Feature, error) {
	feature := &domain.Feature{}

	res := DB.Where("slug = ?", slug).First(feature)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return feature, nil
}

func DeleteFeature(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Feature{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListFeatures() ([]*domain.Feature, error) {
	features := make([]*domain.Feature, 0)

	res := DB.Find(&features)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return features, nil
}
