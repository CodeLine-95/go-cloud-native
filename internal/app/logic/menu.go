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

func MenuResp(c *gin.Context) {
	var params common.SearchRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	var menu models.CloudMenu
	selectFields := structs.ToTags(menu, "json")

	var menuResp []*models.CloudMenu
	err := db.D().Select(selectFields).
		Where("position(concat(?) in concat(menu_name,menu_path,menu_title)) > 0", params.SearchKey).
		Scopes(base.Paginate(params.Page, params.PageSize)).
		Find(&menuResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.PageOK(c, menuResp, len(menuResp), params.Page, params.PageSize, constant.ErrorMsg[constant.Success])
}

func MenuAdd(c *gin.Context) {
	var params common.MenuRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var cloudMenu models.CloudMenu
	cloudMenu.ParseFields(params)
	cloudMenu.SetCreateBy(uint32(auth.UID))
	cloudMenu.CreateTime = uint32(time.Now().Unix())

	res := db.D().Create(cloudMenu)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func MenuEdit(c *gin.Context) {
	var params common.MenuRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var cloudMenu models.CloudMenu
	cloudMenu.ParseFields(params)
	cloudMenu.SetCreateBy(uint32(auth.UID))
	cloudMenu.CreateTime = uint32(time.Now().Unix())

	res := db.D().Updates(&cloudMenu)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func MenuDel(c *gin.Context) {
	var params common.MenuRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	var cloudMenu models.CloudMenu
	cloudMenu.ParseFields(params)
	err := db.D().Delete(cloudMenu).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
