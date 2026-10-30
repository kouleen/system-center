package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryChangeLogPage(ctx context.Context, req *system.SystemChangeLogRequest) (res []*modle.SystemChangeLog, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemChangeLog{}).Where("is_delete = 0")
	if req.GetSchema() != "" {
		query = query.Where("`schema` = ?", req.GetSchema())
	}
	if req.GetTable() != "" {
		query = query.Where("`table` = ?", req.GetTable())
	}
	if req.GetAction() != "" {
		query = query.Where("action = ?", req.GetAction())
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
	if err = query.Count(&total).Error; err != nil || total == 0 {
		return
	}
	query = query.Order("create_time desc,id desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&res).Error; err != nil {
		return
	}
	return
}

func BatchCreateChangeLog(ctx context.Context, entityList []*modle.SystemChangeLog) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(&modle.SystemChangeLog{}).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func BatchDeleteChangeLog(ctx context.Context, idList []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(idList) > 0 {
			if err = tx.Where("id in (?)", idList).Delete(&modle.SystemChangeLog{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
