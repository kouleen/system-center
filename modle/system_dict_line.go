package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemDictLine struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	DictCode   string     `json:"dictCode" gorm:"column:dict_code;not null"`
	DictValue  string     `json:"dictValue" gorm:"column:dict_value;default:''"`
	DictSort   *int32     `json:"dictSort" gorm:"column:dict_sort;default:0"`
	DictType   string     `json:"dictType" gorm:"column:dict_type;not null"`
	ListClass  string     `json:"listClass" gorm:"column:list_class;default:''"`
	Status     *int8      `json:"status" gorm:"column:status;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemDictLine) TableName() string {
	return "system_dict_line"
}

func (p *SystemDictLine) ConvertResp() *system.SystemDictLineResponse {
	return &system.SystemDictLineResponse{
		Id:         p.ID,
		DictCode:   p.DictCode,
		DictValue:  p.DictValue,
		DictSort:   p.DictSort,
		DictType:   p.DictType,
		ListClass:  p.ListClass,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
