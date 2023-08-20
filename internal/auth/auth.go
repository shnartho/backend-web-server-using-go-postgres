package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey exacts an API key from
// the hears of an HTTP request
// Example:
// Authorization: ApiKey {insert api key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
