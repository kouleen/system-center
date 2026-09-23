package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryBulletinPage(ctx context.Context, req *system.SystemBulletinRequest) (resp []*modle.SystemBulletin, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemBulletin{}).Where("is_delete = 0")
	if req.Status != nil {
		query = query.Where("status = ?", req.GetStatus())
	}
	if req.Type != nil {
		query = query.Where("type = ?", req.GetType())
	}
	if req.Title != "" {
		query = query.Where("title like ?", "%"+req.Title+"%")
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
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryBulletinById(ctx context.Context, id int64) (resp *modle.SystemBulletin, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(resp).First(&resp, id).Error; err != nil {
		return
	}
	return
}

func QueryBulletinByIdList(ctx context.Context, ids []int64) (resp []*modle.SystemBulletin, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemBulletin{}).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return
	}
	return
}

func CreateBulletin(ctx context.Context, entity *modle.SystemBulletin) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error
}

func BatchCreateBulletin(ctx context.Context, entityList []*modle.SystemBulletin) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(&modle.SystemBulletin{}).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateBulletin(ctx context.Context, entity *modle.SystemBulletin) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Updates(entity).Error
}

func BatchUpdateBulletin(ctx context.Context, entityList []*modle.SystemBulletin) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range entityList {
			if err = tx.Model(&modle.SystemBulletin{}).Where("id = ?", item.ID).Updates(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
