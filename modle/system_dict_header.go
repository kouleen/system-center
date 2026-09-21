package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemDictHeader struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	DictName   string     `json:"dictName" gorm:"column:dict_name;not null"`
	DictType   string     `json:"dictType" gorm:"column:dict_type;not null"`
	Status     *int8      `json:"status" gorm:"column:status;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemDictHeader) TableName() string {
	return "system_dict_header"
}

func (p *SystemDictHeader) ConvertResp() *system.SystemDictHeaderResponse {
	return &system.SystemDictHeaderResponse{
		Id:         p.ID,
		DictName:   p.DictName,
		DictType:   p.DictType,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
