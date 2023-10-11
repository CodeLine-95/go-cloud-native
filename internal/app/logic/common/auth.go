package common

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/common"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/CodeLine-95/go-cloud-native/tools/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	params := models.LoginRequest{}
	var err error
	if err = c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
	}

	engine := db.D()
	var user models.CloudUser
	has, err := engine.Where("user_name = ?", params.UserName).Get(&user)
	if !has || err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	ok, err := common.CompareHashAndPassword(user.PassWord, params.PassWord)
	if !ok || err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth := jwt.Auth{
		Type: jwt.TypeJWT,
		UID:  user.Id,
		Foo:  utils.RandStringRunes(10),
	}

	token, err := auth.Encode(constant.JwtSignKey)
	if err != nil {
		response.Error(c, constant.ErrorJWT, err, constant.ErrorMsg[constant.ErrorJWT])
		return
	}

	// 更新用户登录信息
	user.LoginIp = c.ClientIP()
	user.LastTime = time.Now().Unix()
	_, _ = engine.Update(user)

	// 写入 Header
	token.SetHeader(c.Writer)

	xlog.Info(traceId.GetLogContext(c, "generate token", logz.F("token", token)))
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
