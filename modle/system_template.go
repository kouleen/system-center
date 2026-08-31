package modle

import "time"

type SystemTemplate struct {
	ID              int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	TemplateCode    string     `json:"templateCode" gorm:"column:template_code;not null;unique"`
	TemplateName    string     `json:"templateName" gorm:"column:template_name;not null"`
	TemplateType    string     `json:"templateType" gorm:"column:template_type;not null"`
	TemplateContent string     `json:"templateContent" gorm:"column:template_content;not null"`
	IsDelete        uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	Remark          string     `json:"remark" gorm:"column:remark;default:''"`
	CreatedBy       int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy       int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime      *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime      *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (s SystemTemplate) TableName() string {
	return "system_template"
}
