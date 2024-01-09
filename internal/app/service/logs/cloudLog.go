package logs

import (
	"github.com/CodeLine-95/go-cloud-native/initial/store/db"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/base"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwtToken"
)

func SaveData(data *models.CloudLog, token jwtToken.Token) error {
	auth := token.Decode(base.JwtSignKey, false)
	if auth != nil {
		var cloudUser models.CloudUser
		if err := db.D().Where("id = ?", auth.UID).Find(&cloudUser).Error; err != nil {
			return err
		}
		data.RequestUser = cloudUser.UserName
	}
	tx := db.D().Begin()
	if err := tx.Table(data.TableName()).Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//func detectTable(db *gorm.DB, c *models.CloudLog) bool {
//	var tableResp string
//	err := db.Raw(fmt.Sprintf("show tables like '%v'", c.TableName())).Scan(&tableResp).Error
//	if err != nil {
//		return false
//	}
//
//	if len(tableResp) == 0 {
//		sql := c.GetCreateSql()
//		if err := db.Exec(sql).Error; err != nil {
//			return false
//		}
//	}
//	return true
//}
