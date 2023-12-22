package base

import (
	"encoding/json"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwtToken"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	TrafficKey = "X-Request-Id"
	JwtSignKey = "__jwt-token-sign__"
)

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

// GetAuth 解析上下文
func GetAuth(c *gin.Context) (auth *jwtToken.Auth, err error) {
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

// Paginate 解析分页和偏移量
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
