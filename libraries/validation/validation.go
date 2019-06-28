package validation

import (
	"fmt"

	"github.com/monzo/terrors"
)

func ErrMissingParam(param string) error {
	return terrors.BadRequest(fmt.Sprintf("missing_param.%s", param), fmt.Sprintf("Param '%s' is required", param), nil)
}

func ErrBadParam(param, message string) error {
	return terrors.BadRequest(fmt.Sprintf("bad_param.%s", param), message, nil)
}

func ErrMissingOneOf(params ...string) error {
	return terrors.BadRequest("missing_param.one_of", fmt.Sprintf("At least one of '%s' is required", params), nil)
}

// AtLeastOne returns true if at least one param is present.
func AtLeastOne(params ...string) bool {
	for _, param := range params {
		if param != "" {
			return true
		}
	}

	return false
}
