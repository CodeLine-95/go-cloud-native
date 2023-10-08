package role

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/resp"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/gin-gonic/gin"
)

// List 角色列表
func List(c *gin.Context) {
	var params models.RoleListRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}
}

// Add 添加角色
func Add(c *gin.Context) {
	var params models.RoleListRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}
}

// Edit 编辑角色
func Edit(c *gin.Context) {
	var params models.RoleListRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}
}

// Del 删除角色
func Del(c *gin.Context) {
	var params models.RoleListRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}
}
