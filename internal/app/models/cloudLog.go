package models

import (
	"encoding/json"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	localTime "github.com/CodeLine-95/go-cloud-native/internal/pkg/time"
	"time"
)

var tableName = "cloud_log_" + time.Now().Format(localTime.DateDayOnly)

type CloudLog struct {
	common.Model
	LogID       string `json:"log_id" gorm:"size:200;not null;uniqueIndex;default:'';comment:日志ID"`
	LogName     string `json:"log_name" gorm:"size:100;not null;index;default:'';comment:日志名称"`
	RequestUrl  string `json:"request_url" gorm:"size:255;not null;index;default:'';comment:请求接口"`
	Method      string `json:"method" gorm:"size:100;not null;index;default:'';comment:请求方法"`
	RequestUser string `json:"request_user" gorm:"size:100;not null;index;default:'';comment:请求用户"`
	ClientIP    string `json:"client_ip" gorm:"size:200;not null;index;default:'';comment:客户端IP"`
	Level       string `json:"Level" gorm:"size:100;not null;index;default:'';comment:级别"`
	AppType     uint32 `json:"app_type" gorm:"not null;index;default:0;comment:应用类型"`
	ParamsData  string `json:"params_data" gorm:"text;default:'';comment:请求参数"`
	common.ModelTime
}

func (c CloudLog) TableName() string {
	postfix := time.Now().Format(localTime.DateDayOnly)
	return "cloud_log_" + postfix
}

func (c *CloudLog) ParseFields(p any) *CloudLog {
	if p == nil {
		return c
	}
	pjson, err := json.Marshal(p)
	if err != nil {
		return c
	}

	err = json.Unmarshal(pjson, c)
	if err != nil {
		return c
	}
	return c
}

//func (c CloudLog) GetCreateSql() string {
//	sql := `CREATE TABLE ` + c.TableName() + ` (
//  id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '唯一编号',
//  log_id varchar(200) NOT NULL COMMENT '日志ID',
//  log_name varchar(100) NOT NULL,
//  request_url varchar(255) NOT NULL,
//  method varchar(100) NOT NULL,
//  request_user varchar(100) NOT NULL,
//  client_ip varchar(200) NOT NULL,
//  level varchar(50) NOT NULL,
//  app_type varchar(50) NOT NULL,
//  params_data text NULL,
//  create_time int(11) unsigned NOT NULL DEFAULT '0',
//  update_time int(11) unsigned NOT NULL DEFAULT '0',
//  PRIMARY KEY (id),
//  KEY log_id (log_id),
//  KEY log_name (log_name),
//  KEY request_url (request_url),
//  KEY method (method),
//  KEY request_user (request_user),
//  KEY client_ip (client_ip),
//  KEY level (level),
//  KEY app_type (log_type),
//  KEY create_time (create_time),
//  KEY update_time (update_time)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台操作日志表';`
//	return sql
//}
