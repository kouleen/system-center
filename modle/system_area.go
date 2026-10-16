package modle

import (
	"time"

	"github.com/kouleen/idl/kitex_gen/system"
)

type SystemArea struct {
	ID            int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	AreaCode      string     `json:"areaCode" gorm:"column:area_code;not null"`
	AreaName      string     `json:"areaName" gorm:"column:area_name;not null"`
	AreaDesc      string     `json:"areaDesc" gorm:"column:area_desc;not null"`
	ParentCode    string     `json:"parentCode" gorm:"column:parent_code;"`
	ParentName    string     `json:"parentName" gorm:"column:parent_name;"`
	RegionCode    string     `json:"regionCode" gorm:"column:region_code;"`
	RegionName    string     `json:"regionName" gorm:"column:region_name;"`
	RegionDesc    string     `json:"regionDesc" gorm:"column:region_desc;"`
	CountryCode   string     `json:"countryCode" gorm:"column:country_code;"`
	CountryName   string     `json:"countryName" gorm:"column:country_name;"`
	CountryDesc   string     `json:"countryDesc" gorm:"column:country_desc;"`
	PhoneCode     string     `json:"phoneCode" gorm:"column:phone_code;"`
	AreaLevelCode string     `json:"areaLevelCode" gorm:"column:area_level_code;"`
	Remark        string     `json:"remark" gorm:"column:remark;default:''"`
	IsDelete      *int8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	CreatedBy     int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy     int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime    *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime    *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (p *SystemArea) TableName() string {
	return "system_area"
}

func (p *SystemArea) ConvertResp() *system.SystemAreaResponse {
	return &system.SystemAreaResponse{
		Id:            p.ID,
		AreaCode:      p.AreaCode,
		AreaName:      p.AreaName,
		AreaDesc:      p.AreaDesc,
		ParentCode:    p.ParentCode,
		ParentName:    p.ParentName,
		RegionCode:    p.RegionCode,
		RegionName:    p.RegionName,
		CountryCode:   p.CountryCode,
		CountryName:   p.CountryName,
		PhoneCode:     p.PhoneCode,
		AreaLevelCode: p.AreaLevelCode,
		Remark:        p.Remark,
		IsDelete:      p.IsDelete,
		CreatedBy:     p.CreatedBy,
		UpdatedBy:     p.UpdatedBy,
		CreateTime:    p.CreateTime.UnixMilli(),
	}
}
