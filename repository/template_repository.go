package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryTemplatePage(ctx context.Context, req *system.SystemTemplateRequest) (list []*modle.SystemTemplate, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("is_delete = ?", 0)
	if req.TemplateCode != "" {
		query = query.Where("template_code = ?", req.TemplateCode)
	}
	if req.TemplateType != "" {
		query = query.Where("template_type = ?", req.TemplateType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}
	if req.TemplateName != "" {
		query = query.Where("template_name like ?", req.TemplateName+"%")
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc,id desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&list).Error; err != nil {
		return
	}
	return
}

func QueryTemplateById(ctx context.Context, id int64) (entity *modle.SystemTemplate, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	return
}

func CreateTemplate(ctx context.Context, entity *modle.SystemTemplate) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Create(entity).Error
}

func UpdateTemplate(ctx context.Context, entity *modle.SystemTemplate) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemTemplate{}).Where("id = ?", entity.ID).Save(entity).Error
}

func DeleteTemplate(ctx context.Context, ids []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("id in (?)", ids).Delete(&modle.SystemTemplate{}).Error; err != nil {
			return err
		}
		return nil
	})
}
