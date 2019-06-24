package dao

import (
	"testing"
	"time"

	"github.com/lolibrary/lolibrary/libraries/idgen"
	"github.com/lolibrary/lolibrary/service.brand/domain"
)

func TestReadBrand(t *testing.T) {
	// TODO: mock database.

	brandID := idgen.RandomTestID()
	imageID := idgen.RandomTestID()
	now := time.Now().UTC()

	testCases := []struct {
		name      string
		database  []*domain.Brand
		id        string
		expected  *domain.Brand
		errPrefix string
	}{
		{
			name: "Read Sole Brand",
			database: []*domain.Brand{
				{
					ID:          brandID,
					Name:        "Angelic Pretty",
					Slug:        "angelic-pretty",
					ShortName:   "ap",
					Description: "Angelic Pretty",
					ImageID:     imageID,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			id: brandID,
			expected: &domain.Brand{
				ID:          brandID,
				Name:        "Angelic Pretty",
				Slug:        "angelic-pretty",
				ShortName:   "ap",
				Description: "Angelic Pretty",
				ImageID:     imageID,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func (t *testing.T) {
			// for _, b := range tc.database {
			// 	err := CreateBrand(b)
			// 	require.NoError(t, err)
			// }
			//
			// brand, err := ReadBrand(tc.id)
			//
			// if tc.errPrefix == "" {
			// 	require.Error(t, err)
			// 	assert.True(t, terrors.Matches(err, tc.errPrefix), "Error has the prefix %v", tc.errPrefix)
			// } else {
			// 	require.NoError(t, err)
			// }
			//
			// assert.Equal(t, tc.expected, brand)
		})
	}

}
