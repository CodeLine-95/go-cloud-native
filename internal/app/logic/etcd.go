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
	"github.com/CodeLine-95/go-cloud-native/tools/pagination"
	"github.com/CodeLine-95/go-cloud-native/tools/structs"
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
	var cloudEtcd models.CloudEtcd
	// 验证roleKey标识，唯一
	var count int64
	err := db.D().Model(cloudEtcd).Where("name = ? and is_delete = ?", params.Name, 0).Count(&count).Error
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
	// etcd 注册
	err = etcdClient.PutService(params.Name, params.Content)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, "etcd put service fail, err: ", logz.F("err", err)))
		response.Error(c, constant.ErrorEtcd, err, constant.ErrorMsg[constant.ErrorEtcd])
		return
	}
	// 入库
	cloudEtcd.ParseFields(params)
	cloudEtcd.SetCreateBy(uint32(auth.UID))
	cloudEtcd.CreateTime = uint32(time.Now().Unix())
	cloudEtcd.IsRegister = 1
	res := db.D().Create(&cloudEtcd)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

func GetService(c *gin.Context) {
	var params common.SearchRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	selectFields := structs.ToTags(models.CloudEtcd{}, "json")
	var cloudEtcdResp []*models.CloudEtcd
	pageList := &pagination.Pagination{
		PageIndex: params.Page,
		PageSize:  params.PageSize,
	}
	err := db.D().Scopes(pagination.Paginate(
		cloudEtcdResp,
		pageList,
		db.D().Select(selectFields).Where("position(? in `name`)", params.SearchKey),
	)).Where("is_delete = ?", 0).Find(&cloudEtcdResp).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	pageList.Rows = cloudEtcdResp
	response.OK(c, pageList, constant.ErrorMsg[constant.Success])
}

func DelService(c *gin.Context) {
	var params common.EtcdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	var cloudEtcd models.CloudEtcd
	cloudEtcd.ParseFields(params)
	err := db.D().Save(cloudEtcd).Error
	if err != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	// 删除 etcd 注册的 key
	err = etcdClient.DelService(params.Name)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, "etcd del service fail, err: ", logz.F("err", err)))
	}
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// 手动注册 / 撤销重新注册
func PutService(c *gin.Context) {
	var params common.EtcdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	var cloudEtcd models.CloudEtcd
	cloudEtcd.ParseFields(params)
	cloudEtcd.IsRegister = 1
	// etcd 注册
	err := etcdClient.PutService(params.Name, params.Content)
	if err != nil {
		xlog.Error(traceId.GetLogContext(c, "etcd put service fail, err: ", logz.F("err", err)))
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	cloudEtcd.IsRegister = 1
	// 更新
	res := db.D().Save(cloudEtcd)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, err, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}

// 服务订阅
func SubscribeService(c *gin.Context) {
	var params common.EtcdRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	var cloudEtcd models.CloudEtcd
	cloudEtcd.ParseFields(params)
	// 更新
	res := db.D().Save(cloudEtcd)
	if res.RowsAffected == 0 || res.Error != nil {
		response.Error(c, constant.ErrorDB, res.Error, constant.ErrorMsg[constant.ErrorDB])
		return
	}
	response.OK(c, nil, constant.ErrorMsg[constant.Success])
}
