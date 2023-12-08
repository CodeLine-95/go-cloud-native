package logic

import (
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/etcd"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/gin-gonic/gin"
	"time"
)

var etcdClient etcd.EtcdClient

func init() {
	etcdClient = etcd.NewClient()
}

func CreatePut(c *gin.Context) {
	var params common.EtcdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}

	// 入库
	var cloudEtcd models.CloudEtcd

	// 验证roleKey标识，唯一
	var count int64
	err := db.D().Model(cloudEtcd).Where("name = ?", params.Name).Count(&count).Error
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

	cloudEtcd.ParseFields(params)
	cloudEtcd.SetCreateBy(uint32(auth.UID))
	cloudEtcd.CreateTime = uint32(time.Now().Unix())

	// etcd 注册
	err = etcdClient.PutService(params.Name, params.Content)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, "etcd put service fail, err: ", logz.F("err", err)))
		cloudEtcd.IsRegister = 1
	}

	res := db.D().Create(&cloudEtcd)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}

	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func GetService(c *gin.Context) {

}
