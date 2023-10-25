package middleware

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var name string
		nameInfo, err := c.Request.Cookie("userName")
		if err == nil && nameInfo.Value != "" {
			name = nameInfo.Value
		}
		xlog.Info(traceId.GetLogContext(c, "JWTLogin", logz.F("name", name)))

		if userName, ok := c.Get(constant.UserName); ok {
			xlog.Info(traceId.GetLogContext(c, "JWTLogin", logz.F("name", name), logz.F("userName", userName)))
			c.Next()
		} else {
			accept := c.Request.Header.Get("Accept")
			if strings.Index(accept, "html") > -1 {
				c.Abort()
				return
			} else {
				// 获取 token
				token := jwt.GetToken(c.Request, "")
				// 验证token非空
				if token == "" {
					response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
					return
				}
				// token验证是否失效
				auth := token.Decode(base.JwtSignKey, false)
				if auth == nil {
					response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
					return
				}
				// 设置到上下文
				c.Set("auth", auth)
				c.Next()
			}
		}

	}
}
