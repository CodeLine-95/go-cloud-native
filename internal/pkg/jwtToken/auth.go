package jwtToken

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func init() {
	RegisterValidFunc(TypeJWT, defaultJWTValidFunc)
}

// RegisterValidFunc 注册校验函数
func RegisterValidFunc(authType string, validFunc ValidFunc) {
	validFuncs[authType] = validFunc
}

func defaultJWTValidFunc(a *Auth) error {
	if a.UID == 0 {
		return errors.New("uid is empty")
	}
	return nil
}

// Valid 校验Auth
func (a *Auth) Valid() error {
	if a == nil {
		return errors.New("auth is empty")
	}
	if a.ExpiresAt.Unix() < time.Now().Unix() {
		return errors.New("auth is expired")
	}
	if valid, ok := validFuncs[a.Subject]; ok {
		return valid(a)
	}
	return errors.New("unknown auth type")
}

// Encode 将Auth编码成Token
func (a *Auth) Encode(sign string) (Token, error) {
	if a.ExpiresAt == nil || a.ExpiresAt.Unix() <= 0 {
		a.ExpiresAt = jwt.NewNumericDate(time.Unix(time.Now().Unix()+DefaultDuration, 0))
	}
	a.Subject = TypeJWT
	// 验证 Auth
	if err := a.Valid(); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	t, err := token.SignedString([]byte(sign))
	return Token(t), err
}

func (a *Auth) String() string {
	data, _ := json.Marshal(a)
	return string(data)
}
