package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemRole struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	RoleName   string     `json:"roleName" gorm:"column:role_name;not null;default:''"`
	RoleSort   *int32     `json:"roleSort" gorm:"column:role_sort;not null"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemRole) TableName() string {
	return "system_role"
}

func (p *SystemRole) ConvertResp() *system.SystemRoleResponse {
	return &system.SystemRoleResponse{
		Id:         p.ID,
		RoleName:   p.RoleName,
		RoleSort:   p.RoleSort,
		Status:     p.Status,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
