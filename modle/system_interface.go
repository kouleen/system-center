package modle

import "time"

type SystemInterface struct {
	ID            int64      `json:"id,string" gorm:"column:id;primary_key;not null"`
	RequestPath   string     `json:"requestPath" gorm:"column:request_path;not null"`
	InterfaceName string     `json:"interfaceName" gorm:"column:interface_name;not null"`
	MethodType    uint8      `json:"methodType" gorm:"column:method_type;not null;default:1"`
	MethodName    string     `json:"methodName" gorm:"column:method_name;not null"`
	ParamTypes    string     `json:"paramTypes" gorm:"column:param_types;default:''"`
	Version       string     `json:"version" gorm:"column:version;not null"`
	Status        uint8      `json:"status" gorm:"column:status;not null;default:1"`
	IsDelete      uint8      `json:"isDelete" gorm:"column:is_delete;not null;default:0"`
	Remark        string     `json:"remark" gorm:"column:remark;default:''"`
	CreatedBy     int64      `json:"createdBy,string" gorm:"column:created_by;not null;default:-1"`
	UpdatedBy     int64      `json:"updatedBy,string" gorm:"column:updated_by;not null;default:-1"`
	CreateTime    *time.Time `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	UpdateTime    *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (SystemInterface) TableName() string {
	return "system_interface"
}
