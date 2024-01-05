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
	"gorm.io/gorm/clause"
)

func GetUserList(c *gin.Context) {
	var params common.SearchRequest
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	selectFields := structs.ToTags(models.CloudUser{}, "json")
	var userResp []*models.CloudUser
	err := db.D().Select(selectFields).
		Where("position(concat(?) in concat(user_name)) > 0", params.SearchKey).
		Scopes(base.Paginate(params.Page, params.PageSize)).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "admin"}, Desc: true}).
		Find(&userResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	response.PageOK(c, userResp, len(userResp), params.Page, params.PageSize, constant.ErrorMsg[constant.Success])
}

func GetUserInfo(c *gin.Context) {
	auth, err := base.GetAuth(c)
	if err != nil {
		response.Error(c, constant.ErrorNotLogin, err, constant.ErrorMsg[constant.ErrorNotLogin])
		return
	}

	var userResp models.GetCloudUserAndRole
	res := db.D().Where("id = ?", auth.UID).Find(&userResp)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, userResp, constant.ErrorMsg[constant.Success])
}
