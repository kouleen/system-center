package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemMenu struct {
	ID         int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	MenuName   string     `json:"menuName" gorm:"column:menu_name;not null"`
	ParentId   int64      `json:"parentId,string" gorm:"column:parent_id;default:0;not null"`
	OrderNum   *int32     `json:"orderNum" gorm:"column:order_num;default:0;not null"`
	Path       string     `json:"path" gorm:"column:path;default:''"`
	Component  string     `json:"component" gorm:"column:component;default:''"`
	Query      string     `json:"query" gorm:"column:query;default:''"`
	RouteName  string     `json:"routeName" gorm:"column:route_name;default:''"`
	IsFrame    *int8      `json:"isFrame" gorm:"column:is_frame;not null;default:0"`
	IsCache    *int8      `json:"isCache" gorm:"column:is_cache;not null;default:1"`
	MenuType   *int8      `json:"menuType" gorm:"column:menu_type;not null;default:1"`
	Visible    *int8      `json:"visible" gorm:"column:visible;not null;default:1"`
	Status     *int8      `json:"status" gorm:"column:status;not null;default:1"`
	Perms      string     `json:"perms" gorm:"column:perms;default:''"`
	Icon       string     `json:"icon" gorm:"column:icon;default:'#'"`
	Remark     string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete   *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy  int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy  int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemMenu) TableName() string {
	return "system_menu"
}

func (p *SystemMenu) ConvertResp() *system.SystemMenuResponse {
	return &system.SystemMenuResponse{
		Id:         p.ID,
		MenuName:   p.MenuName,
		ParentId:   p.ParentId,
		OrderNum:   p.OrderNum,
		Path:       p.Path,
		Component:  p.Component,
		Query:      p.Query,
		RouteName:  p.RouteName,
		IsFrame:    p.IsFrame,
		IsCache:    p.IsCache,
		MenuType:   p.MenuType,
		Visible:    p.Visible,
		Status:     p.Status,
		Perms:      p.Perms,
		Icon:       p.Icon,
		IsDelete:   p.IsDelete,
		CreatedBy:  p.CreatedBy,
		UpdatedBy:  p.UpdatedBy,
		CreateTime: p.CreateTime.UnixMilli(),
	}
}
