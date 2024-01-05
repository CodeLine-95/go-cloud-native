package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/time"
	"github.com/CodeLine-95/go-cloud-native/tools/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetLogList(c *gin.Context) {
	var params common.LogSearchReqest
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	startTime, _ := time.ParseTime(params.StartTime)
	endTime, _ := time.ParseTime(params.EndTime)

	selectFields := structs.ToTags(models.CloudLog{}, "json")
	var logResp []*models.CloudLog
	err := db.D().Select(selectFields).
		Where("position(concat(?) in concat(log_name)) > 0", params.SearchKey).
		Where("create_time >= ? and create_time <= ?", startTime, endTime).
		Scopes(base.Paginate(params.Page, params.PageSize)).
		Order(clause.OrderByColumn{Column: clause.Column{Name: "admin"}, Desc: true}).
		Find(&logResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.PageOK(c, logResp, len(logResp), params.Page, params.PageSize, constant.ErrorMsg[constant.Success])
}
