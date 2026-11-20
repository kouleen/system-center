package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemRoleArea struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	RoleId     int64      `json:"roleId,string" gorm:"column:role_id;not null"`
	AreaId     int64      `json:"areaId,string" gorm:"column:area_id;not null"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemRoleArea) TableName() string {
	return "system_role_area"
}

func (p *SystemRoleArea) ConvertResp() *system.SystemRoleAreaResponse {
	return &system.SystemRoleAreaResponse{
		Id:         p.ID,
		RoleId:     p.RoleId,
		AreaId:     p.AreaId,
		Remark:     p.Remark,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
