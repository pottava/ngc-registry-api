package registry

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims used at NGC authentication
type Claims struct {
	Subject   string `json:"sub,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	ClientID  string `json:"authClientId,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	ID        string `json:"jti,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
}

// ParseJWT parses token to JWT & its claims
func ParseJWT(token string) (*jwt.Token, *Claims, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if !strings.EqualFold(token.Method.Alg(), jwt.SigningMethodRS256.Alg()) {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return "without valid key", nil
	})
	if err != nil {
		if !strings.EqualFold(err.Error(), jwt.ErrInvalidKeyType.Error()) {
			return nil, nil, fmt.Errorf("Couldn't handle: %s", err.Error())
		}
	}
	claims, ok := parsed.Claims.(*Claims)
	if !ok {
		return nil, nil, fmt.Errorf("Couldn't parse as standard claims")
	}
	if err = claims.Valid(); err != nil {
		return nil, nil, err
	}
	return parsed, claims, nil
}

func cookie(session string) *http.Cookie {
	return &http.Cookie{
		Name:     "SessionToken",
		Value:    session,
		Path:     "/",
		Domain:   "api.ngc.nvidia.com",
		MaxAge:   36000,
		HttpOnly: true,
		Secure:   true,
	}
}

// Valid returns true if the jwt is active
func (c *Claims) Valid() error {
	err := jwt.ValidationError{}
	now := jwt.TimeFunc().Unix()

	if now >= c.ExpiresAt {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		err.Inner = fmt.Errorf("token is expired by %v", delta)
		err.Errors |= jwt.ValidationErrorExpired
	}
	if now <= c.IssuedAt {
		err.Inner = fmt.Errorf("Token used before issued")
		err.Errors |= jwt.ValidationErrorIssuedAt
	}
	if now <= c.NotBefore {
		err.Inner = fmt.Errorf("token is not valid yet")
		err.Errors |= jwt.ValidationErrorNotValidYet
	}
	if err.Errors == 0 {
		return nil
	}
	return err
}
