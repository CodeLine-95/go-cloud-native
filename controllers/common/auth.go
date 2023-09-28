package common

import (
	"github.com/CodeLine-95/go-cloud-native/common/constant"
	"github.com/CodeLine-95/go-cloud-native/pkg/jwt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/resp"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/services"
	"github.com/CodeLine-95/go-cloud-native/services/models"
	"github.com/CodeLine-95/go-cloud-native/store/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	params := services.LoginRequest{}
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
	_, err = engine.Where("user_name = ?", params.UserName).Get(&user)
	if err != nil {
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
		Type:     jwt.TypeJWT,
		UID:      int64(user.Id),
		UserName: user.UserName,
		Exp:      0,
	}

	token, err := auth.Encode(jwt.JWTSIGN)
	if err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorJWT], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorJWT,
			ErrMsg:  constant.ErrorMsg[constant.ErrorJWT],
		}, nil)
		return
	}

	token.SetHeader(c.Writer)

	xlog.Info(traceId.GetLogContext(c, "generate token", logz.F("token", token)))
	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, gin.H{
		"token": token,
	})
}
