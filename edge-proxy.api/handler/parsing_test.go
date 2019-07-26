package handler

import (
	"testing"

	"github.com/monzo/terrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsePath(t *testing.T) {
	testCases := []struct{
		name string
		path string
		expectedService, expectedPath string
		errPrefix string
	}{
		{
			name: "Blank Path",
			path: "",
			errPrefix: "not_found.service",
		},
		{
			name: "Single Slash",
			path: "/",
			errPrefix: "not_found.service",
		},
		{
			name: "Service without Path",
			path: "/images",
			expectedService: "s-api-images.lolibrary",
			expectedPath: "/",
		},
		{
			name: "Service with Slash",
			path: "/images/",
			expectedService: "s-api-images.lolibrary.lolibrary",
			expectedPath: "/",
		},
		{
			name: "Service with Path",
			path: "/images/upload/presigned-url",
			expectedService: "s-api-images.lolibrary.lolibrary",
			expectedPath: "/upload/presigned-url",
		},
		{
			name: "Bad Service",
			path: "/test service/upload/presigned-url",
			errPrefix: "bad_request.service",
		},
		{
			name: "Service Ending with Dash",
			path: "/test-service-0123-/upload/presigned-url",
			errPrefix: "bad_request.service",
		},
		{
			name: "Service with Numbers and Dashes",
			path: "/test-service-123/upload/presigned-url",
			expectedService: "s-api-test-service-123.lolibrary",
			expectedPath: "/upload/presigned-url",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc // re-bind because of parallel goroutine
			t.Parallel()

			service, path, err := parsePath(tc.path)

			if tc.errPrefix != "" {
				require.Error(t, err)
				assert.True(t, terrors.PrefixMatches(err, tc.errPrefix))
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tc.expectedService, service)
			assert.Equal(t, tc.expectedPath, path)
		})
	}
}