package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemLoginLog struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	Username   string     `json:"username" gorm:"column:username;not null"`
	IP         string     `json:"ip" gorm:"column:ip;default:''"`
	Location   string     `json:"location" gorm:"column:location;default:''"`
	OS         string     `json:"os" gorm:"column:os;default:''"`
	Token      string     `json:"token" gorm:"column:token;not null"`
	Browser    string     `json:"browser" gorm:"column:browser;default:''"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:1"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (*SystemLoginLog) TableName() string {
	return "system_login_log"
}

func (p *SystemLoginLog) ConvertResp() *system.SystemLoginLogResponse {
	return &system.SystemLoginLogResponse{
		Id:         p.ID,
		Username:   p.Username,
		Ip:         p.IP,
		Location:   p.Location,
		Os:         p.OS,
		Token:      p.Token,
		Browser:    p.Browser,
		Status:     p.Status,
		Remark:     p.Remark,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
