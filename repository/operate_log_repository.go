package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryOperateLogPage(ctx context.Context, req *system.SystemOperateLogRequest) (resp []*modle.SystemOperateLog, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemOperateLog{}).Where("is_delete = ?", 0)
	if req.GetModule() != "" {
		query.Where("module = ?", req.GetModule())
	}
	if req.GetBusiness() != "" {
		query.Where("business = ?", req.GetBusiness())
	}
	if req.Type != "" {
		query.Where("type = ?", req.GetType())
	}
	if req.Status != nil {
		query = query.Where("status = ?", req.GetStatus())
	}
	if req.Params != nil {
		if req.GetParams().GetBeginTime() != "" && req.GetParams().GetEndTime() != "" {
			startUTC, endUTC, err := utils.DateToLocalRange(req.GetParams().GetBeginTime(), req.GetParams().GetEndTime())
			if err != nil {
				return nil, 0, err
			}
			query = query.Where("create_time between ? and ?", startUTC, endUTC)
		}
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func CreateOperateLog(ctx context.Context, entity *modle.SystemOperateLog) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(entity).Create(entity).Error
}

func BatchDeleteOperateLog(ctx context.Context, idList []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(idList) > 0 {
			if err = tx.Where("id in (?)", idList).Delete(&modle.SystemOperateLog{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
