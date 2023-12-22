package middleware

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/initial/store/policy"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwtToken"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Policy() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 角色
		// 获取 token
		token := jwtToken.GetToken(c.Request, "")
		// 验证token非空
		if token == "" {
			response.Error(c, http.StatusOK, nil, constant.ErrorMsg[constant.ErrorNotLogin])
			return
		}
		// token验证是否失效
		auth := token.Decode(base.JwtSignKey, false)
		if auth == nil {
			response.Error(c, http.StatusOK, nil, constant.ErrorMsg[constant.ErrorNotLogin])
			return
		}

		//获取路径
		obj := c.Request.URL.Path
		// 获取方法
		act := c.Request.Method
		// 角色
		var userRole models.CloudUserRole
		err := db.D().Where("uid = ?", auth.UID).Find(&userRole).Error
		if err != nil {
			response.Error(c, http.StatusOK, nil, constant.ErrorMsg[constant.ErrorNotLogin])
			return
		}
		sub := userRole.RoleId

		// 验证当前访问策略
		if ok := policy.Casbin().Checked(sub, obj, act); !ok {
			response.Error(c, http.StatusOK, nil, constant.ErrorMsg[constant.ErrorNotLogin])
			return
		}
	}
}
