package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

func UserRole(c *gin.Context) {
	var params common.UserRoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var cloudUserRole models.CloudUserRole
	cloudUserRole.ParseFields(params)

	// 验证是否是更新
	if params.IsUpdate == 1 {
		cloudUserRole.SetUpdateBy(uint32(auth.UID))
		cloudUserRole.UpdateTime = uint32(time.Now().Unix())
	} else {
		cloudUserRole.SetCreateBy(uint32(auth.UID))
		cloudUserRole.CreateTime = uint32(time.Now().Unix())
	}

	res := db.D().Save(&cloudUserRole)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func RoleMenu(c *gin.Context) {
	var params common.RoleMenuRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var cloudRoleMenu models.CloudRoleMenu
	db.D().Where("role_id = ?", params.RoleId).Delete(cloudRoleMenu)

	cloudRoleMenu.ParseFields(params)
	cloudRoleMenu.SetCreateBy(uint32(auth.UID))
	cloudRoleMenu.CreateTime = uint32(time.Now().Unix())
	cloudRoleMenu.SetUpdateBy(uint32(auth.UID))
	cloudRoleMenu.UpdateTime = uint32(time.Now().Unix())

	res := db.D().Create(&cloudRoleMenu)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
