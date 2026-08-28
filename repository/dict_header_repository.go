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
		return nil, 0, err
	}
	query = query.Order("create_time desc")
	i := (req.GetPageRequest().GetCurrent() - 1) * req.GetPageRequest().GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetPageRequest().GetSize())).Find(&systemDictHeaderList).Error; err != nil {
		return nil, 0, err
	}
	return systemDictHeaderList, total, nil
}
