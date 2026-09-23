package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemEducation struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Title      string     `json:"title" gorm:"column:title;not null"`
	Hits       *int8      `json:"hits" gorm:"column:hits;not null"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemEducation) TableName() string {
	return "system_education"
}

func (p *SystemEducation) ConvertResp() *system.SystemEducationResponse {
	var updateTime int64
	if p.UpdateTime != nil {
		updateTime = p.UpdateTime.Unix()
	}
	return &system.SystemEducationResponse{
		Id:         p.ID,
		Title:      p.Title,
		Hits:       p.Hits,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
		UpdateTime: updateTime,
	}
}
