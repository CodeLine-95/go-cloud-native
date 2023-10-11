package role

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/resp"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/gin-gonic/gin"
	"time"
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
	var params models.RoleAddRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}

	cloudRole := &models.CloudRole{
		Name:      params.Name,
		Remark:    params.Remark,
		RulesIds:  params.RulesIds,
		CreatTime: time.Now().Unix(),
	}
	engine := db.D()
	cnt, err := engine.Insert(cloudRole)
	if cnt == 0 || err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorDB], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorDB,
			ErrMsg:  constant.ErrorMsg[constant.ErrorDB],
		}, nil)
		return
	}

	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, nil)
}

// Edit 编辑角色
func Edit(c *gin.Context) {
	var params models.RoleEditRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}

	cloudRole := &models.CloudRole{
		Id:         params.Id,
		Name:       params.Name,
		Remark:     params.Remark,
		RulesIds:   params.RulesIds,
		Status:     params.Status,
		UpdateTime: time.Now().Unix(),
	}
	engine := db.D()
	cnt, err := engine.Update(cloudRole)
	if cnt == 0 || err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorDB], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorDB,
			ErrMsg:  constant.ErrorMsg[constant.ErrorDB],
		}, nil)
		return
	}

	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, nil)
}

// Del 删除角色
func Del(c *gin.Context) {
	var params models.CloudRole
	if err := c.ShouldBindJSON(&params); err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorParams], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}

	engine := db.D()
	cnt, err := engine.Delete(params)
	if cnt == 0 || err != nil {
		xlog.Info(traceId.GetLogContext(c, constant.ErrorMsg[constant.ErrorDB], logz.F("err", err)))
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorDB,
			ErrMsg:  constant.ErrorMsg[constant.ErrorDB],
		}, nil)
		return
	}

	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, nil)
}
