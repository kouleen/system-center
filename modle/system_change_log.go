package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemChangeLog struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Schema     string     `json:"schema" gorm:"column:schema;not null"`
	Table      string     `json:"table" gorm:"column:table;not null"`
	Action     string     `json:"action" gorm:"column:action;not null"`
	Before     string     `json:"before" gorm:"column:before;default:''"`
	After      string     `json:"after" gorm:"column:after;default:''"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (*SystemChangeLog) TableName() string {
	return "system_change_log"
}

func (p *SystemChangeLog) ConvertResp() *system.SystemChangeLogResponse {
	return &system.SystemChangeLogResponse{
		Id:         p.ID,
		Schema:     p.Schema,
		Table:      p.Table,
		Action:     p.Action,
		Before:     p.Before,
		After:      p.After,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
