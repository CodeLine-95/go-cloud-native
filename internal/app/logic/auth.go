package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwtToken"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/ip"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/CodeLine-95/go-cloud-native/tools/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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
	jwtExpTime := viper.GetInt("jwt.expireTime")
	auth := jwtToken.Auth{
		utils.RandStringRunes(10),
		int64(userResp.Id),
		int64(userResp.RoleId),
		int64(userResp.Admin),
		jwtToken.AuthExtend{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+int64(jwtExpTime), 0)),
		},
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
