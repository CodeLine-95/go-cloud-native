package logic

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

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
