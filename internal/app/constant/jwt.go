package constant

import (
	"encoding/json"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	JwtSignKey = "__jwt-token-sign__"
)

// GetAuth 解析上下文
func GetAuth(c *gin.Context) (auth *jwt.Auth, err error) {
	authAny, _ := c.Get("auth")
	authJson, err := json.Marshal(authAny)
	if err != nil {
		return
	}
	err = json.Unmarshal(authJson, &auth)
	if err != nil {
		return
	}
	return
}
