package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryEducationPage(ctx context.Context, req *system.SystemEducationRequest) (resp []*modle.SystemEducation, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemEducation)).Where("is_delete = 0")
	if req.GetTitle() != "" {
		query.Where("title LIKE ?", "%"+req.GetTitle()+"%")
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

func QueryEducationById(ctx context.Context, id int64) (resp *modle.SystemEducation, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemEducation)).First(&resp, id).Error; err != nil {
		return
	}
	return
}

func QueryEducationByIdList(ctx context.Context, ids []int64) (resp []*modle.SystemEducation, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemEducation)).Where("id in (?)").Find(&resp, ids).Error; err != nil {
	}
	return
}

func CreateEducation(ctx context.Context, entity *modle.SystemEducation) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error
}

func BatchCrateEducation(ctx context.Context, entityList []*modle.SystemEducation) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(new(modle.SystemEducation)).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateEducation(ctx context.Context, entity *modle.SystemEducation) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Save(entity).Error
}

func BatchUpdateEducation(ctx context.Context, entityList []*modle.SystemEducation) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, education := range entityList {
			if err = tx.Model(new(modle.SystemEducation)).Where("id = ?", education.ID).Save(education).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
