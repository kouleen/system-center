package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryDictLinePage(ctx context.Context, req *system.SystemDictLineRequest) (systemDictLines []*modle.SystemDictLine, total int64, err error) {
	query := getDictLineQuery(ctx, req)
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("dict_sort")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&systemDictLines).Error; err != nil {
		return
	}
	return
}

func getDictLineQuery(ctx context.Context, req *system.SystemDictLineRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictLine{}).Where("is_delete = ?", 0)
	if req.DictType != "" {
		query = query.Where("dict_type = ?", req.DictType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.DictCode != "" {
		query = query.Where("dict_code like ?", "%"+req.DictCode+"%")
	}
	if req.DictValue != "" {
		query = query.Where("dict_value like ?", "%"+req.DictValue+"%")
	}
	return query
}

func QueryDictLineList(ctx context.Context, req *system.SystemDictLineRequest) (systemDictLines []*modle.SystemDictLine, err error) {
	query := getDictLineQuery(ctx, req)
	query = query.Order("create_time desc")
	if err = query.Find(&systemDictLines).Error; err != nil {
		return nil, err
	}
	return systemDictLines, nil
}

func QueryDictLineById(ctx context.Context, id int64) (systemDictLine *modle.SystemDictLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictLine{}).First(&systemDictLine, id).Error; err != nil {
		return
	}
	return
}

func CreateDictLine(ctx context.Context, entity *modle.SystemDictLine) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemDictLine{}).Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDictLine(ctx context.Context, entity *modle.SystemDictLine) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemDictLine{}).Where("id = ?", entity.ID).Save(entity).Error; err != nil {
		return err
	}
	return nil
}
