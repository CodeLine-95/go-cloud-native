package jwt

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	TypeJWT = "jwt"
	JWTSIGN = "__jwt-token-sign__"
)

type ValidFunc func(c *Auth) error

var validFuncs = make(map[string]ValidFunc)

type Auth struct {
	Type     string `json:"type,omitempty"`     // 认证方式
	UID      int64  `json:"uid,omitempty"`      // UID
	UserName string `json:"userName,omitempty"` //用户名
	Exp      int64  `json:"exp,omitempty"`      //过期时间
}

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
	if a.Exp < time.Now().Unix() {
		return errors.New("auth is expired")
	}
	if valid, ok := validFuncs[a.Type]; ok {
		return valid(a)
	}
	return errors.New("unknown auth type")
}

// Encode 将Auth编码成Token
func (a *Auth) Encode(sign string) (Token, error) {
	if a.Exp <= 0 {
		a.Exp = time.Now().Unix() + DefaultDuration
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	t, err := token.SignedString([]byte(sign))
	return Token(t), err
}

func (a *Auth) String() string {
	data, _ := json.Marshal(a)
	return string(data)
}
