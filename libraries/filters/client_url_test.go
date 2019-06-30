package filters

import (
	"testing"

	"github.com/monzo/typhon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceToHost(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Basic Service",
			input:  "service.foo",
			output: "s-bar",
		},
		{
			name:   "Basic API Service",
			input:  "service.api.foo",
			output: "s-api-foo",
		},
		{
			name:   "Service with Hyphens",
			input:  "service.this-is-a-test",
			output: "s-this-is-a-test",
		},
		{
			name:   "Edge Proxy Internal",
			input:  "edge-proxy.internal",
			output: "edge-proxy-internal",
		},
		{
			name:   "Edge Proxy External",
			input:  "edge-proxy.external",
			output: "edge-proxy-external",
		},
		{
			name:   "Unknown New Type",
			input:  "unknown.foo.bar",
			output: "unknown-foo-bar",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.output, serviceToHost(tc.input))
		})
	}
}

func TestClientURLFilter(t *testing.T) {
	testCases := []struct {
		name string
		req  typhon.Request

		scheme, host, method, path string
	}{
		{
			name:   "Single Level Service",
			req:    typhon.NewRequest(nil, "GET", "/service.foo/list", []byte("{}")),
			host:   "s-foo",
			method: "GET",
			path:   "/list",
		},
		{
			name:   "Two Level Service",
			req:    typhon.NewRequest(nil, "GET", "/service.foo.bar/list", []byte("{}")),
			host:   "s-foo-bar",
			method: "GET",
			path:   "/list",
		},
		{
			name:   "API Service",
			req:    typhon.NewRequest(nil, "GET", "/service.api.foo/test/abc/foo", []byte("{}")),
			host:   "s-api-foo",
			method: "GET",
			path:   "/test/abc/foo",
		},
		{
			name:   "POST Request",
			req:    typhon.NewRequest(nil, "POST", "/service.foo/test/abc/foo", []byte("{}")),
			host:   "s-foo",
			method: "POST",
			path:   "/test/abc/foo",
		},
		{
			name:   "PUT Request",
			req:    typhon.NewRequest(nil, "PUT", "/service.foo/test/abc/foo", []byte("{}")),
			host:   "s-foo",
			method: "PUT",
			path:   "/test/abc/foo",
		},
		{
			name:   "DELETE Request",
			req:    typhon.NewRequest(nil, "DELETE", "/service.foo/test/abc/foo", []byte("{}")),
			host:   "s-foo",
			method: "DELETE",
			path:   "/test/abc/foo",
		},
		{
			name:   "Unrelated URL",
			req:    typhon.NewRequest(nil, "GET", "https://test.example/foo", nil),
			scheme: "https",
			host:   "test.example",
			method: "GET",
			path:   "/foo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_ = ClientURLFilter(tc.req, func(req typhon.Request) typhon.Response {
				require.NotNil(t, req)

				assert.Equal(t, tc.host, req.URL.Host)
				assert.Equal(t, tc.path, req.URL.Path)

				if tc.scheme != "" {
					assert.Equal(t, tc.scheme, req.URL.Scheme)
				} else {
					assert.Equal(t, "http", req.URL.Scheme)
				}

				assert.Equal(t, tc.method, req.Method)
				assert.Equal(t, tc.host, req.Host)

				return req.Response(nil)
			})
		})
	}
}
