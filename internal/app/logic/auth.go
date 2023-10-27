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
	var userResp models.GetCloudUserAndRole
	err = engine.Where("user_name = ?", params.UserName).Find(&userResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	ok, err := base.CompareHashAndPassword(userResp.PassWord, params.PassWord)
	if !ok || err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth := jwt.Auth{
		Type:    jwt.TypeJWT,
		UID:     int64(userResp.Id),
		Foo:     utils.RandStringRunes(10),
		RoleID:  int64(userResp.RoleId),
		IsAdmin: int64(userResp.Admin),
	}

	token, err := auth.Encode(base.JwtSignKey)
	if err != nil {
		response.Error(c, constant.ErrorJWT, err, constant.ErrorMsg[constant.ErrorJWT])
		return
	}

	// 更新用户登录信息
	var user models.CloudUser
	user.ParseFields(userResp)
	user.LoginIp = ip.ClientIP(c)
	user.LastTime = uint32(time.Now().Unix())
	err = engine.Save(user).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	xlog.Info(traceId.GetLogContext(c, "generate token", logz.F("token", token)))
	response.OK(c, gin.H{
		"token": token,
	}, constant.ErrorMsg[constant.Success])
}
