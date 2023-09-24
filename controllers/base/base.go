package base

import (
	"github.com/CodeLine-95/go-cloud-native/services/params"
	"github.com/gin-gonic/gin"
)

// ApiInterface 通用API
type ApiInterface interface {
	Login(c *gin.Context) // 用户登录
	UpLoad(c *gin.Context)
}

type CommonApi struct {
}

func (api CommonApi) Login(c *gin.Context) {

	param := &params.LoginParams{}
	paramErr := c.ShouldBindJSON(&param)
	if paramErr != nil {
		return
	}
	c.JSON(200, gin.H{})

	return
}
