package dao

import (
	"context"
	"testing"

	"github.com/lolibrary/lolibrary/service.item/domain"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestReadAndCreateItem(t *testing.T) {
	testCases := []struct {
		name string
		item *domain.Item
	}{
		{
			name: "Regular Item",
			item: &domain.Item{
				ID:            "b012831d-eec5-4931-a2f2-6ab3e9268ad2",
				Slug:          "ap-test-fake-product",
				BrandID:       "9e9ab949-026b-491d-8617-362f6a9cbc91",
				CategoryID:    "b462eb8f-077b-45ed-be89-0e0ec1701e46",
				UserID:        "472cee53-79d2-4060-89bf-2ef69e7c6a92",
				ImageID:       "094cd4d2-8c30-4087-be59-e604f9f5ab33",
				PublisherID:   "cb24b502-8323-4058-8523-5d4f6f29da72",
				EnglishName:   "Test English Name",
				ForeignName:   "テスト",
				ProductNumber: "12345⭐ABC",
				Currency:      "jpy",
				Price:         "34000",
			},
		},
		{
			name: "Regular Item with Metadata",
			item: &domain.Item{
				ID:            "a6a9d842-8669-4d63-9de1-413b701f2f8a",
				Slug:          "ap-test-fake-product-2",
				BrandID:       "9e9ab949-026b-491d-8617-362f6a9cbc91",
				CategoryID:    "b462eb8f-077b-45ed-be89-0e0ec1701e46",
				UserID:        "472cee53-79d2-4060-89bf-2ef69e7c6a92",
				ImageID:       "094cd4d2-8c30-4087-be59-e604f9f5ab33",
				PublisherID:   "cb24b502-8323-4058-8523-5d4f6f29da72",
				EnglishName:   "Test English Name",
				ForeignName:   "テスト2",
				ProductNumber: "12346⭐ABC",
				Currency:      "jpy",
				Price:         "34000",
				Metadata: map[string]string{
					"original_id": "124",
					"original_slug": "test-fake-product-2",
				},
			},
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := CreateItem(ctx, tc.item)
			require.NoError(t, err)

			item1, err := ReadItem(ctx, tc.item.ID)
			require.NoError(t, err)

			item2, err := ReadItemBySlug(ctx, tc.item.Slug)
			require.NoError(t, err)

			assert.DeepEqual(t, tc.item, item1)
			assert.DeepEqual(t, tc.item, item2)
		})
	}
}
