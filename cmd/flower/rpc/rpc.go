// Package rpc performs RPC calls for a user, returning responses via typhon.
//
// The package should be initialised with:
// rpc.SetInternalEdgeProxy()
// rpc.SetExternalEdgeProxy()
//
// You may need to `kubectl port-forward service/edge-proxy-internal 8080:80`.
package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	"github.com/monzo/slog"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

var (
	internalEdgeProxyURL = ""
	externalEdgeProxyURL = ""

	errRequest = fmt.Errorf("")
)

func SetInternalEdgeProxy(url string) {
	internalEdgeProxyURL = url
}

func SetExternalEdgeProxy(url string) {
	externalEdgeProxyURL = url
}

// Request sends an RPC call through the internal edge proxy.
func InternalRequest(ctx context.Context, method, path, body string) (string, error) {
	if internalEdgeProxyURL == "" {
		return "", terrors.New("bad_request.env_not_set", "Environment has not been set.", nil)
	}

	switch method {
	case "GET", "POST", "PUT", "DELETE", "PATCH":
	case "get", "post", "put", "delete", "patch":
	default:
		return "", terrors.New("bad_request.bad_param.method", "Method must be a valid HTTP method.", nil)
	}

	b := []byte(body)

	req := typhon.NewRequest(ctx, method, internalEdgeProxyURL+path, nil)
	if n, err := req.Write(b); err != nil || n != len(b) {
		slog.Error(ctx, "Response not fully written. Aborting Request.")
		return "", err
	}

	rsp := req.Send().Response()
	if rsp.Error != nil {
		slog.Error(ctx, "Error sending request via typhon: %v", rsp.Error)
		return "", rsp.Error
	}

	// pretty-print the returned body
	buf, err := rsp.BodyBytes(true)
	if err != nil {
		slog.Error(ctx, "Error getting body contents: %v", err)
		return "", terrors.Wrap(err, nil)
	}

	// We got an error. Format it and remove the stack.
	if rsp.Header.Get("Terror") == "1" {
		return handleError(ctx, buf)
	}

	responseBody, err := prettyjson.Format(buf)
	if err != nil {
		slog.Error(ctx, "Error formatting response JSON: %v", err)
		return "", terrors.Wrap(err, nil)
	}

	return string(responseBody), nil
}

// Request sends an RPC call through the internal edge proxy.
// It returns a pretty-printed body on success, or an error on failure.
func ExternalRequest(path, body string) (string, error) {
	return "", terrors.New("not_implemented", "Not yet implemented", nil)
}

func handleError(ctx context.Context, buf []byte) (string, error) {
	errCheck := make(map[string]interface{})
	if err := json.Unmarshal(buf, &errCheck); err != nil {
		slog.Error(ctx, "Failed to unmarshal error body: %v", err)
		return string(buf), terrors.Wrap(err, nil)
	}

	// if there is a "stack" key, delete it.
	delete(errCheck, "stack")

	// pretty print it
	responseBody, err := prettyjson.Marshal(errCheck)
	if err != nil {
		slog.Error(ctx, "Unable to pretty-print error body: %v", err)
		return "", terrors.Wrap(err, nil)
	}

	return string(responseBody), errRequest
}
