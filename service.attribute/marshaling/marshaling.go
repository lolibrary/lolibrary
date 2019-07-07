package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.attribute/domain"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
)

func AttributeToProto(attribute *domain.Attribute) *attributeproto.Attribute {
	if attribute == nil {
		return nil
	}

	return &attributeproto.Attribute{
		Id:        attribute.ID,
		Slug:      attribute.Slug,
		Name:      attribute.Name,
		CreatedAt: util.TimeToProto(attribute.CreatedAt),
		UpdatedAt: util.TimeToProto(attribute.UpdatedAt),
	}
}

func AttributesToProto(attributes []*domain.Attribute) []*attributeproto.Attribute {
	protos := make([]*attributeproto.Attribute, 0, len(attributes))
	for _, attribute := range attributes {
		protos = append(protos, AttributeToProto(attribute))
	}

	return protos
}
