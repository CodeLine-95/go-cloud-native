package middleware

import (
	"github.com/CodeLine-95/go-cloud-native/common/constant"
	"github.com/CodeLine-95/go-cloud-native/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
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
				token := jwt.GetToken(c.Request, "user")
				// 验证token非空
				if token == "" {
					c.AbortWithStatusJSON(http.StatusOK, constant.Response{
						ErrNo: constant.ErrorNotLogin,
						Msg:   constant.ErrorMsg[constant.ErrorNotLogin],
						Data:  nil,
					})
					c.Abort()
					return
				}
				// token验证是否失效
				auth := token.Decode(string(token), true)
				if auth == nil {
					c.AbortWithStatusJSON(http.StatusOK, constant.Response{
						ErrNo: constant.ErrorNotLogin,
						Msg:   constant.ErrorMsg[constant.ErrorNotLogin],
						Data:  nil,
					})
					c.Abort()
					return
				}
				c.Next()
			}
		}

	}
}
