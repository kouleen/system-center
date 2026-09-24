package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryLoginLogPage(ctx context.Context, req *system.SystemLoginLogRequest) (resp []*modle.SystemLoginLog, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemLoginLog{}).Where("is_delete = ?", 0)
	if req.Status != nil {
		query = query.Where("status = ?", req.GetStatus())
	}
	if req.GetUsername() != "" {
		query = query.Where("username LIKE ?", req.GetUsername()+"%")
	}
	if req.GetIp() != "" {
		query = query.Where("method_type like ?", req.GetIp()+"%")
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

func QueryLoginLogByIdList(ctx context.Context, idList []int64) (resp []*modle.SystemLoginLog, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemLoginLog{}).Where("id in (?)", idList).Find(&resp).Error; err != nil {
		return
	}
	return
}

func CreateLoginLog(ctx context.Context, entity *modle.SystemLoginLog) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemLoginLog{}).Create(entity).Error
}

func BatchDeleteLoginLog(ctx context.Context, idList []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(idList) > 0 {
			if err = tx.Where("id in (?)", idList).Delete(&modle.SystemLoginLog{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
