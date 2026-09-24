package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
)

func QueryTemplatePage(ctx context.Context, req *system.SystemTemplateRequest) (list []*modle.SystemTemplate, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("is_deleted = ?", 0)
	if req.TemplateCode == "" {
		query = query.Where("template_code = ?", req.TemplateCode)
	}
	if req.TemplateType == "" {
		query = query.Where("template_type = ?", req.TemplateType)
	}
	if req.TemplateName == "" {
		query = query.Where("template_name like ?", req.TemplateName+"%")
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&list).Error; err != nil {
		return
	}
	return
}

func QueryTemplateById(ctx context.Context, id int64) (entity *modle.SystemTemplate, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("id = ?", id).First(&entity).Error; err != nil {
		return
	}
	return
}

func CreateTemplate(ctx context.Context, entity *modle.SystemTemplate) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Create(entity).Error; err != nil {
		return
	}
	return
}

func UpdateTemplate(ctx context.Context, entity *modle.SystemTemplate) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("id = ?", entity.ID).Save(entity).Error; err != nil {
		return
	}
	return
}
