package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemBulletin struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Title      string     `json:"title" gorm:"column:title;not null"`
	Type       *int8      `json:"type" gorm:"column:type;not null"`
	Status     *int8      `json:"status" gorm:"column:status;default:1"`
	Content    string     `json:"content" gorm:"column:content;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemBulletin) TableName() string {
	return "system_bulletin"
}

func (p *SystemBulletin) ConvertResp() *system.SystemBulletinResponse {
	return &system.SystemBulletinResponse{
		Id:         p.ID,
		Title:      p.Title,
		Type:       p.Type,
		Status:     p.Status,
		Content:    p.Content,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
