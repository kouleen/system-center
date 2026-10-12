package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemOperateLog struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Module     string     `json:"module" gorm:"column:module;not null"`
	Business   string     `json:"business" gorm:"column:business;not null"`
	Type       string     `json:"type" gorm:"column:type;not null"`
	IP         string     `json:"ip" gorm:"column:ip;default:''"`
	Location   string     `json:"location" gorm:"column:location;default:''"`
	OS         string     `json:"os" gorm:"column:os;default:''"`
	Browser    string     `json:"browser" gorm:"column:browser;default:''"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (*SystemOperateLog) TableName() string {
	return "system_operate_log"
}

func (p *SystemOperateLog) ConvertResp() *system.SystemOperateLogResponse {
	return &system.SystemOperateLogResponse{
		Id:         p.ID,
		Module:     p.Module,
		Business:   p.Business,
		Type:       p.Type,
		Ip:         p.IP,
		Location:   p.Location,
		Os:         p.OS,
		Browser:    p.Browser,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
