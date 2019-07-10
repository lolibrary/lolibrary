package validation

import (
	"fmt"
	"regexp"

	"github.com/monzo/terrors"
	uuid "github.com/nu7hatch/gouuid"
)

var reSlug = regexp.MustCompile(`^[a-z0-9][a-z0-9\-]+[a-z0-9]$`)

func ErrMissingParam(param string) error {
	return terrors.BadRequest(fmt.Sprintf("missing_param.%s", param), fmt.Sprintf("Param '%s' is required", param), nil)
}

func ErrBadParam(param, message string) error {
	return terrors.BadRequest(fmt.Sprintf("bad_param.%s", param), message, nil)
}

func ErrBadSlug(param, slug string) error {
	return ErrBadParam(param, fmt.Sprintf("'%s' is not a valid slug", slug))
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

// UUID checks if a given UUID is valid.
func UUID(id string) bool {
	_, err := uuid.ParseHex(id)
	return err == nil
}

// Slug checks if a slug is valid.
func Slug(slug string) bool {
	return reSlug.MatchString(slug)
}
