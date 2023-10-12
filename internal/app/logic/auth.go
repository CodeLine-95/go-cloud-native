package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/ip"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/CodeLine-95/go-cloud-native/tools/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	params := common.LoginRequest{}
	var err error
	if err = c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	engine := db.D()
	var user models.CloudUser
	err = engine.Where("user_name = ?", params.UserName).Find(&user).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	ok, err := base.CompareHashAndPassword(user.PassWord, params.PassWord)
	if !ok || err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth := jwt.Auth{
		Type: jwt.TypeJWT,
		UID:  int64(user.Id),
		Foo:  utils.RandStringRunes(10),
	}

	token, err := auth.Encode(constant.JwtSignKey)
	if err != nil {
		response.Error(c, constant.ErrorJWT, err, constant.ErrorMsg[constant.ErrorJWT])
		return
	}

	// 更新用户登录信息
	user.LoginIp = ip.ClientIP(c)
	user.LastTime = uint32(time.Now().Unix())
	err = engine.Save(user).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	// 写入 Header
	token.SetHeader(c.Writer)

	xlog.Info(traceId.GetLogContext(c, "generate token", logz.F("token", token)))
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
