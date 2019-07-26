package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.attributes/domain"
	attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"
)

func ProtoToAttributes(protos []*attributeproto.Attribute) []*domain.Attribute {
	attributes := make([]*domain.Attribute, 0, len(protos))
	for _, attribute := range protos {
		attributes = append(attributes, ProtoToAttribute(attribute))
	}

	return attributes
}

func ProtoToAttribute(attribute *attributeproto.Attribute) *domain.Attribute {
	if attribute == nil {
		return nil
	}

	return &domain.Attribute{
		Slug: attribute.Slug,
		Name: attribute.Name,
	}
}
