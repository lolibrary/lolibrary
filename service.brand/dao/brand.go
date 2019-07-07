package dao

import (
	"github.com/lolibrary/lolibrary/libraries/database"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	"github.com/monzo/terrors"
)

func CreateBrand(brand *domain.Brand) error {
	res := DB.Create(brand)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func UpdateBrand(brand *domain.Brand) error {
	res := DB.Update(brand)
	if res.Error != nil {
		if err := database.DuplicateRecord(res.Error); err != nil {
			return err
		}

		if res.RecordNotFound() {
			return terrors.NotFound("brand", "Brand not found", nil)
		}

		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ReadBrand(id string) (*domain.Brand, error) {
	brand := &domain.Brand{}

	res := DB.Where("id = ?", id).First(brand)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return brand, nil
}

func ReadBrandBySlug(slug string) (*domain.Brand, error) {
	brand := &domain.Brand{}

	res := DB.Where("slug = ?", slug).First(brand)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return brand, nil
}

func ReadBrandByShortName(shortName string) (*domain.Brand, error) {
	brand := &domain.Brand{}

	res := DB.Where("short_name = ?", shortName).First(brand)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return brand, nil
}

func DeleteBrand(id string) error {
	res := DB.Where("id = ?", id).Delete(&domain.Brand{})
	if res.Error != nil {
		return terrors.Wrap(res.Error, nil)
	}

	return nil
}

func ListBrands() ([]*domain.Brand, error) {
	brands := make([]*domain.Brand, 0)

	res := DB.Find(&brands)
	if res.Error != nil {
		if res.RecordNotFound() {
			return nil, nil
		}

		return nil, terrors.Wrap(res.Error, nil)
	}

	return brands, nil
}
