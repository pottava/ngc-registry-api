package lib

import (
	"errors"
	"strings"

	"github.com/go-openapi/swag"
)

// Principal principal
type Principal struct {
	Session string
}

// RequestToPrincipal returns principal from token string in the request 'Authorization' header
func RequestToPrincipal(token string) (*Principal, error) {
	token = strings.Replace(strings.Replace(token, "Bearer ", "", 1), "Token ", "", 1)
	if swag.IsZero(token) {
		return nil, errors.New("Invalid token")
	}
	return &Principal{Session: token}, nil
}
