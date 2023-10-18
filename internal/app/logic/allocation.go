package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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

	if params.RoleId == 0 || params.MenuIds == "" {
		response.Error(c, constant.ErrorParams, nil, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	// 每次创建时，清空当前角色的权限关系，重新插入
	db.D().Where("role_id = ?", params.RoleId).Delete(models.CloudRoleMenu{})

	var cloudRoleMenu []*models.CloudRoleMenu
	menuIdList := strings.Split(params.MenuIds, ",")
	for _, val := range menuIdList {
		menuId, _ := strconv.Atoi(val)
		cloudRoleMenu = append(cloudRoleMenu, &models.CloudRoleMenu{
			RoleId: params.RoleId,
			MenuId: uint32(menuId),
			ControlBy: common.ControlBy{
				CreateBy: uint32(auth.UID),
				UpdateBy: uint32(auth.UID),
			},
			ModelTime: common.ModelTime{
				CreateTime: uint32(time.Now().Unix()),
				UpdateTime: uint32(time.Now().Unix()),
			},
		})
	}
	res := db.D().Create(&cloudRoleMenu)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
