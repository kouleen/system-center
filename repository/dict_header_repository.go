package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
)

func QueryDictHeaderPage(ctx context.Context, req *system.SystemDictHeaderRequest) (systemDictHeaderList []modle.SystemDictHeader, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).Where("is_delete = ?", 0)
	if req.DictType != "" {
		query = query.Where("dict_type = ?", req.DictType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.DictName != "" {
		query = query.Where("dict_name like ?", "%"+req.DictName+"%")
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&systemDictHeaderList).Error; err != nil {
		return
	}
	return systemDictHeaderList, total, nil
}

func QueryDictHeaderList(ctx context.Context, req *system.SystemDictHeaderRequest) (systemDictHeaderList []modle.SystemDictHeader, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).Where("is_delete = ?", 0)
	if req.DictType != "" {
		query = query.Where("dict_type = ?", req.DictType)
	}
	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}
	if req.DictName != "" {
		query = query.Where("dict_name like ?", "%"+req.DictName+"%")
	}
	query = query.Order("create_time desc")
	if err = query.Find(&systemDictHeaderList).Error; err != nil {
		return
	}
	return systemDictHeaderList, nil
}

func QueryDictHeaderById(ctx context.Context, id int64) (systemDictHeader *modle.SystemDictHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&systemDictHeader, id).Error; err != nil {
		return
	}
	return
}

func CreateDictHeader(ctx context.Context, entity *modle.SystemDictHeader) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error; err != nil {
		return
	}
	return nil
}

func UpdateDictHeader(ctx context.Context, entity *modle.SystemDictHeader) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Updates(entity).Error; err != nil {
		return
	}
	return nil
}
