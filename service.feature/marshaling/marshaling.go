package marshaling

import (
	"github.com/lolibrary/lolibrary/libraries/util"
	"github.com/lolibrary/lolibrary/service.feature/domain"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
)

func FeatureToProto(feature *domain.Feature) *featureproto.Feature {
	if feature == nil {
		return nil
	}

	return &featureproto.Feature{
		Id:        feature.ID,
		Slug:      feature.Slug,
		Name:      feature.Name,
		CreatedAt: util.TimeToProto(feature.CreatedAt),
		UpdatedAt: util.TimeToProto(feature.UpdatedAt),
	}
}

func FeaturesToProto(features []*domain.Feature) []*featureproto.Feature {
	protos := make([]*featureproto.Feature, 0, len(features))
	for _, feature := range features {
		protos = append(protos, FeatureToProto(feature))
	}

	return protos
}
