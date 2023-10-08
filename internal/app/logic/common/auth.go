package common

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/resp"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/CodeLine-95/go-cloud-native/tools/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(c *gin.Context) {
	params := models.LoginRequest{}
	var err error
	if err = c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}

	engine := db.Grp(constant.CloudNative)
	var user models.CloudUser
	has, err := engine.Where("user_name = ?", params.UserName).Get(&user)
	if !has || err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorDB], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorDB,
			ErrMsg:  constant.ErrorMsg[constant.ErrorDB],
		}, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(params.PassWord))
	if err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
		return
	}

	auth := jwt.Auth{
		Type: jwt.TypeJWT,
		UID:  user.Id,
		Foo:  utils.RandStringRunes(10),
	}

	token, err := auth.Encode(constant.JwtSignKey)
	if err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorJWT], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorJWT,
			ErrMsg:  constant.ErrorMsg[constant.ErrorJWT],
		}, nil)
		return
	}

	// 更新用户登录信息
	user.LoginIp = c.ClientIP()
	user.LastTime = time.Now().Unix()
	_, _ = engine.Update(user)

	// 写入 Header
	token.SetHeader(c.Writer)

	xlog.Info(traceId.GetLogContext(c, "generate token", logz.F("token", token)))
	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, nil)
}
