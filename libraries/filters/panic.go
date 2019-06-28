package filters

import (
	"fmt"
	"runtime/debug"

	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
)

// PanicFilter catches any panics that happen while serving.
func PanicFilter(req typhon.Request, svc typhon.Service) (rsp typhon.Response) {
	defer func() {
		if p := recover(); p != nil {
			trace := debug.Stack()
			// TODO: logging.
			slogParams := map[string]string{
				"error": fmt.Sprintf("%v", p),
				"trace": string(trace),
			}

			err := terrors.InternalService("panic", fmt.Sprintf("panic: handling %v: %v", req, string(trace)), slogParams)
			rsp = typhon.Response{Error: err}
		}
	}()

	rsp = svc(req)

	return
}
