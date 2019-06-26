package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
)

// BrandToProto marshals a domain object into a protobuf object.
func BrandToProto(brand *domain.Brand) *brandproto.Brand {
	if brand == nil {
		return nil
	}

	return &brandproto.Brand{
		Id:          brand.ID,
		Slug:        brand.Slug,
		ShortName:   brand.ShortName,
		Name:        brand.Name,
		Description: brand.Description,
		ImageId:     brand.ImageID,
		CreatedAt:   util.TimeToProto(brand.CreatedAt),
		UpdatedAt:   util.TimeToProto(brand.UpdatedAt),
	}
}

// BrandsToProto marshals a slice of domain objects into protobuf objects.
func BrandsToProto(brands []*domain.Brand) []*brandproto.Brand {
	protos := make([]*brandproto.Brand, 0, len(brands))
	for _, brand := range brands {
		protos = append(protos, BrandToProto(brand))
	}

	return protos
}
