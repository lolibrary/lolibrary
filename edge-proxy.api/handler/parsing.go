package handler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/monzo/terrors"
)

var re = regexp.MustCompile(`^[a-z0-9][a-z0-9\-]{0,28}[a-z0-9]$`)

const serviceFormat = "s-api-%s.lolibrary.svc.cluster.local"

func parsePath(path string) (string, string, error) {
	switch path {
	case "", "/":
		return "", "", terrors.NotFound("service", "No endpoint was given", nil)
	}

	parts := strings.SplitN(strings.TrimPrefix(path, "/"), "/", 2)
	if len(parts) != 2 {
		parts = append(parts, "")
	}

	// validate the service is alpha-with-dashes only.
	if !re.MatchString(parts[0]) {
		return "", "", terrors.BadRequest("service", "Endpoint is invalid", map[string]string{
			"service": parts[0],
		})
	}

	return fmt.Sprintf(serviceFormat, parts[0]), "/" + parts[1], nil
}
