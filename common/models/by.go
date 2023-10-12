package models

type ControlBy struct {
	CreateBy uint32 `json:"create_by" gorm:"not null;index;default:0;comment:创建者: 填写用户ID"`
	UpdateBy uint32 `json:"update_by" gorm:"not null;index;default:0;comment:更新者: 填写用户ID"`
}

// SetCreateBy 设置创建人ID
func (e *ControlBy) SetCreateBy(createBy uint32) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置更新人ID
func (e *ControlBy) SetUpdateBy(updateBy uint32) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id uint32 `json:"id" gorm:"primaryKey;autoIncrement;comment:唯一编号"`
}

type ModelTime struct {
	CreateTime uint32 `json:"create_time" gorm:"autoCreateTime;not null;index;default:0;comment:创建时间"`
	UpdateTime uint32 `json:"update_time" gorm:"autoUpdateTime;not null;index;default:0;comment:更新时间"`
}
