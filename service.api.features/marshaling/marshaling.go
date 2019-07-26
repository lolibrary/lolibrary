package marshaling

import (
	"github.com/lolibrary/lolibrary/service.api.features/domain"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
)

func ProtoToFeatures(protos []*featureproto.Feature) []*domain.Feature {
	features := make([]*domain.Feature, 0, len(protos))
	for _, feature := range protos {
		features = append(features, ProtoToFeature(feature))
	}

	return features
}

func ProtoToFeature(feature *featureproto.Feature) *domain.Feature {
	if feature == nil {
		return nil
	}

	return &domain.Feature{
		Slug: feature.Slug,
		Name: feature.Name,
	}
}
