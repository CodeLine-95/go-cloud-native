package models

type ControlBy struct {
	CreateBy int64 `json:"create_by" xorm:"int(11) unsigned notnull index default(0) comment('创建者: 填写用户ID')"`
	UpdateBy int64 `json:"update_by" xorm:"int(11) unsigned notnull index default(0) comment('更新者: 填写用户ID')"`
}

// SetCreateBy 设置创建人ID
func (e *ControlBy) SetCreateBy(createBy int64) {
	e.CreateBy = createBy
}

// SetUpdateBy 设置更新人ID
func (e *ControlBy) SetUpdateBy(updateBy int64) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int64 `json:"id" xorm:"pk unsigned comment('主键编码') version"`
}

type ModelTime struct {
	CreatTime  int64 `json:"creat_time" xorm:"int(11) notnull index default(0) comment('创建时间')"`
	UpdateTime int64 `json:"update_time" xorm:"int(11) notnull index default(0) comment('更新时间')"`
}
