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
	"fmt"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/lolibrary/lolibrary/libraries/filters"
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
// Both return values from this function should be printed to stdout if not empty/nil.
func InternalRequest(ctx context.Context, method, path, body string) (string, error) {
	req, err := internalRequest(ctx, method, path, body)
	if err != nil {
		return "", err
	}

	if internalEdgeProxyURL == "" {
		return "", terrors.New("bad_request.env_not_set", "Environment has not been set.", nil)
	}

	client := typhon.Client.
		Filter(filters.EdgeProxyFilter(internalEdgeProxyURL)).
		Filter(typhon.ErrorFilter)

	rsp := req.SendVia(client).Response()
	if rsp.Error != nil {
		return parseError(rsp)
	}

	return parseResponse(rsp)
}

// internalRequest creates a typhon request from options.
func internalRequest(ctx context.Context, method, path, body string) (typhon.Request, error) {
	req := typhon.NewRequest(ctx, method, path, nil)
	if _, err := req.Write([]byte(body)); err != nil {
		return typhon.Request{}, err
	}

	return req, nil
}

// parseResponse returns pretty-printed JSON from responses.
func parseResponse(rsp typhon.Response) (string, error) {
	// pretty-print the returned body
	buf, err := rsp.BodyBytes(true)
	if err != nil {
		return "", terrors.Wrap(err, nil)
	}

	responseBody, err := prettyjson.Format(buf)
	if err != nil {
		return "", terrors.Wrap(err, nil)
	}

	return string(responseBody), nil
}

// parseError returns a representation of any errors in the response.
func parseError(rsp typhon.Response) (string, error) {
	if terr, ok := rsp.Error.(*terrors.Error); ok {
		payload := map[string]interface{}{"message": terr.Message, "code": terr.Code}
		if len(terr.Params) > 0 {
			payload["params"] = terr.Params
		}

		if os.Getenv("DEBUG") == "1" {
			payload["stack"] = terr.StackFrames
		}

		b, err := prettyjson.Marshal(payload)
		if err != nil {
			return "", err
		}

		return string(b), terr
	} else {
		// if we have a body, return that verbatim.
		if b, err := rsp.BodyBytes(true); err == nil {
			return string(b), rsp.Error
		}

		return "", rsp.Error
	}
}
