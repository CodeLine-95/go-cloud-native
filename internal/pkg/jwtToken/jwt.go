package jwtToken

import (
	"encoding/json"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

// GetToken 从请求中获取jwt Token
func GetToken(r *http.Request, cookieName string) Token {
	token := r.Header.Get("X-Auth")
	if token == "" {
		cookie, err := r.Cookie(cookieName)
		if err == nil && cookie.Value != "" {
			token = cookie.Value
		}
	}
	return Token(token)
}

// Decode 将Token解码成Auth结构体， verify为true表示进行，校验失败则返回nil
func (t Token) Decode(sign string, verify bool) *Auth {
	jwtClaims := &AuthExtend{}
	claims := &Auth{}
	parser := &jwt.Parser{}
	if verify {
		parser = jwt.NewParser(jwt.WithoutClaimsValidation())
	}
	token, err := parser.ParseWithClaims(string(t), jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(sign), nil
	})
	if err != nil {
		logz.Error("jwtToken token decode", logz.F("error", err.Error()))
		return nil
	}
	if token == nil || !token.Valid {
		return nil
	}

	jsonClaims, _ := json.Marshal(jwtClaims)
	_ = json.Unmarshal(jsonClaims, claims)

	return claims
}

// SetCookie 将jwt Token保存到cookie中
func (t Token) SetCookie(w http.ResponseWriter, cookieName string) {
	w.Header().Set("Set-Cookie", fmt.Sprintf("%s=%s", cookieName, string(t)))
}

// SetHeader 将jwt Token保存到请求返回的X-Auth头部
func (t Token) SetHeader(w http.ResponseWriter) {
	w.Header().Set("X-Auth", string(t))
}

// String jwtToken Token转成字符串
func (t Token) String() string {
	return string(t)
}
