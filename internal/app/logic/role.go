package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/tools/structs"
	"github.com/gin-gonic/gin"
	"time"
)

func RoleResp(c *gin.Context) {
	var params common.SearchRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	var role models.CloudRole
	selectFields := structs.ToTags(role, "json")

	var roleResp []*models.CloudRole
	err := db.D().Select(selectFields).
		Where("position(concat(?) in concat(role_key,role_name)) > 0", params.SearchKey).
		Scopes(base.Paginate(params.Page, params.PageSize)).
		Find(&roleResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.PageOK(c, roleResp, len(roleResp), params.Page, params.PageSize, constant.ErrorMsg[constant.Success])
}

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	var cloudRole models.CloudRole

	// 验证roleKey标识，唯一
	var count int64
	err := db.D().Model(cloudRole).Where("role_key = ?", params.RoleKey).Count(&count).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	if count > 0 {
		response.Error(c, constant.ErrorDBRecordExist, nil, constant.ErrorMsg[constant.ErrorDBRecordExist])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	cloudRole.ParseFields(params)
	cloudRole.SetCreateBy(uint32(auth.UID))
	cloudRole.CreateTime = uint32(time.Now().Unix())

	res := db.D().Create(&cloudRole)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// RoleEdit 编辑角色
func RoleEdit(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var cloudRole models.CloudRole

	cloudRole.ParseFields(params)
	cloudRole.SetUpdateBy(uint32(auth.UID))
	cloudRole.UpdateTime = uint32(time.Now().Unix())

	res := db.D().Updates(&cloudRole)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// RoleDel 删除角色
func RoleDel(c *gin.Context) {
	var params common.RoleRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	var cloudRole models.CloudRole
	cloudRole.ParseFields(params)
	err := db.D().Delete(cloudRole).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
