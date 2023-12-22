package jwtToken

import "github.com/golang-jwt/jwt/v5"

const (
	TypeJWT         = "jwt"
	DefaultDuration = int64(2 * 3600)
)

type (
	Token string
	Auth  struct {
		Foo     string `json:"foo"`
		UID     int64  `json:"uid"`
		RoleID  int64  `json:"role_id"`
		IsAdmin int64  `json:"is_admin"`
		AuthExtend
	}
	ValidFunc  func(c *Auth) error
	AuthExtend jwt.RegisteredClaims
)

var validFuncs = make(map[string]ValidFunc)

// GetExpirationTime implements the Claims interface.
func (c AuthExtend) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.ExpiresAt, nil
}

// GetNotBefore implements the Claims interface.
func (c AuthExtend) GetNotBefore() (*jwt.NumericDate, error) {
	return c.NotBefore, nil
}

// GetIssuedAt implements the Claims interface.
func (c AuthExtend) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.IssuedAt, nil
}

// GetAudience implements the Claims interface.
func (c AuthExtend) GetAudience() (jwt.ClaimStrings, error) {
	return c.Audience, nil
}

// GetIssuer implements the Claims interface.
func (c AuthExtend) GetIssuer() (string, error) {
	return c.Issuer, nil
}

// GetSubject implements the Claims interface.
func (c AuthExtend) GetSubject() (string, error) {
	return c.Subject, nil
}
