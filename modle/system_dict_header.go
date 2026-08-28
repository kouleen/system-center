package modle

import "time"

type SystemDictHeaderCreate struct {
	ID         *int64     `json:"id,string" gorm:"column:id;primary_key;not null"`
	DictName   *string    `json:"dictName" gorm:"column:dict_name;not null" binding:"required" validate:"required"`
	DictType   *string    `json:"dictType" gorm:"column:dict_type;not null" binding:"required" validate:"required"`
	Status     *uint8     `json:"status" gorm:"column:status;default:1" binding:"required" validate:"required"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''" binding:"required"`
	IsDelete   uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type SystemDictHeader struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	DictName   string     `json:"dictName" gorm:"column:dict_name;not null"`
	DictType   string     `json:"dictType" gorm:"column:dict_type;not null"`
	Status     uint8      `json:"status" gorm:"column:status;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (SystemDictHeader) TableName() string {
	return "system_dict_header"
}
