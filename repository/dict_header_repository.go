package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryDictHeaderPage(ctx context.Context, req *system.SystemDictHeaderRequest) (systemDictHeaderList []modle.SystemDictHeader, total int64, err error) {
	query := getDictHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&systemDictHeaderList).Error; err != nil {
		return
	}
	return
}

func getDictHeaderQuery(ctx context.Context, req *system.SystemDictHeaderRequest) *gorm.DB {
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
	return query
}

func QueryDictHeaderList(ctx context.Context, req *system.SystemDictHeaderRequest) (systemDictHeaderList []modle.SystemDictHeader, err error) {
	query := getDictHeaderQuery(ctx, req)
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	query = query.Order("create_time desc")
	if err = query.Find(&systemDictHeaderList).Error; err != nil {
		return
	}
	return systemDictHeaderList, nil
}

func QueryDictHeaderById(ctx context.Context, id int64) (systemDictHeader *modle.SystemDictHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).First(&systemDictHeader, id).Error; err != nil {
		return
	}
	return
}

func QueryDictHeaderByIdList(ctx context.Context, ids []int64) (systemDictHeaderList []*modle.SystemDictHeader, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).Where("id in (?)", ids).Find(&systemDictHeaderList).Error; err != nil {
		return
	}
	return
}

func CreateDictHeader(ctx context.Context, entity *modle.SystemDictHeader) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).Create(entity).Error
}

func BatchCreateDictHeader(ctx context.Context, entityList []*modle.SystemDictHeader) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(&modle.SystemDictHeader{}).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateDictHeader(ctx context.Context, entity *modle.SystemDictHeader) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemDictHeader{}).Where("id = ?", entity.ID).Save(entity).Error
}

func BatchUpdateDictHeader(ctx context.Context, entityList []*modle.SystemDictHeader) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, header := range entityList {
			if err = tx.Model(&modle.SystemDictHeader{}).Where("id = ?", header.ID).Save(header).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
