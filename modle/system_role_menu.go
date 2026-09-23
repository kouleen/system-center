package modle

import "time"

type SystemRoleMenu struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	RoleId     int64      `json:"roleId,string" gorm:"column:role_id;not null"`
	MenuId     int64      `json:"menuId,string" gorm:"column:menu_id;not null"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (SystemRoleMenu) TableName() string {
	return "system_role_menu"
}
