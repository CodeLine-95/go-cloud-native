package role

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

// List 角色列表
func List(c *gin.Context) {
	var params common.RoleListRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
}

// Add 添加角色
func Add(c *gin.Context) {
	var params common.RoleAddRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := constant.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	cloudRole := &models.CloudRole{
		RoleName:   params.Name,
		RoleRemark: params.Remark,
		RoleKey:    params.Key,
		RoleSort:   params.Sort,
		ControlBy: common.ControlBy{
			CreateBy: auth.UID,
		},
		ModelTime: common.ModelTime{
			CreatTime: time.Now().Unix(),
		},
	}
	engine := db.D()
	cnt, err := engine.Insert(cloudRole)
	if cnt == 0 || err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// Edit 编辑角色
func Edit(c *gin.Context) {
	var params common.RoleEditRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := constant.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	cloudRole := &models.CloudRole{
		RoleId:     params.Id,
		RoleName:   params.Name,
		RoleRemark: params.Remark,
		RoleKey:    params.Key,
		RoleSort:   params.Sort,
		Status:     params.Status,
		ControlBy: common.ControlBy{
			UpdateBy: auth.UID,
		},
		ModelTime: common.ModelTime{
			UpdateTime: time.Now().Unix(),
		},
	}
	engine := db.D()
	cnt, err := engine.Update(cloudRole)
	if cnt == 0 || err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// Del 删除角色
func Del(c *gin.Context) {
	var params models.CloudRole
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	engine := db.D()
	cnt, err := engine.Delete(params)
	if cnt == 0 || err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
