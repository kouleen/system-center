package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
)

func QueryDictLinePage(ctx context.Context, req *system.SystemDictLineRequest) (systemDictLines []*modle.SystemDictLine, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Where("is_delete = ?", 0)
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
	if err = query.Count(&total).Error; err != nil {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&systemDictLines).Error; err != nil {
		return
	}
	return systemDictLines, total, nil
}

func QueryDictLineList(ctx context.Context, req *system.SystemDictLineRequest) (systemDictLines []*modle.SystemDictLine, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Where("is_delete = ?", 0)
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
	query = query.Order("create_time desc")
	if err = query.Find(&systemDictLines).Error; err != nil {
		return nil, err
	}
	return systemDictLines, nil
}

func QueryDictLineById(ctx context.Context, id int64) (systemDictLine *modle.SystemDictLine, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&systemDictLine, id).Error; err != nil {
		return
	}
	return
}

func CreateDictLine(ctx context.Context, entity *modle.SystemDictLine) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDictLine(ctx context.Context, entity *modle.SystemDictLine) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Updates(entity).Error; err != nil {
		return err
	}
	return nil
}
