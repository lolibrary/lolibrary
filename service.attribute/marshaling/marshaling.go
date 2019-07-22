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

func AttributeValueToProto(value *domain.AttributeValue) *attributeproto.AttributeValue {
	if value == nil {
		return nil
	}

	return &attributeproto.AttributeValue{
		Id:          value.ID,
		AttributeId: value.AttributeID,
		ItemId:      value.ItemID,
		Value:       value.Value,
	}
}

func AttributeValuesToProto(values []*domain.AttributeValue) []*attributeproto.AttributeValue {
	protos := make([]*attributeproto.AttributeValue, 0, len(values))
	for _, val := range values {
		protos = append(protos, AttributeValueToProto(val))
	}

	return protos
}
