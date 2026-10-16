package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryAreaPage(ctx context.Context, req *system.SystemAreaRequest) (resp []*modle.SystemArea, total int64, err error) {
	query := getAreaQuery(ctx, req)
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
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func getAreaQuery(ctx context.Context, req *system.SystemAreaRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemArea{}).Where("is_delete = 0")
	if req.GetAreaCode() != "" {
		query = query.Where("area_code = ?", req.GetAreaCode())
	}
	if req.GetRegionCode() != "" {
		query = query.Where("region_code = ?", req.GetRegionCode())
	}
	if req.GetParentCode() != "" {
		query = query.Where("parent_code = ?", req.GetParentCode())
	}
	if req.GetCountryCode() != "" {
		query = query.Where("country_code = ?", req.GetCountryCode())
	}
	if req.GetAreaLevelCode() != "" {
		query = query.Where("area_level_code = ?", req.GetAreaLevelCode())
	}
	if req.GetPhoneCode() != "" {
		query = query.Where("phone_code = ?", req.GetPhoneCode())
	}
	if req.GetAreaName() != "" {
		query = query.Where("area_name like ?", "%"+req.GetAreaName()+"%")
	}
	if req.GetParentName() != "" {
		query = query.Where("parent_name like ?", "%"+req.GetParentName()+"%")
	}
	if req.GetRegionName() != "" {
		query = query.Where("region_name like ?", "%"+req.GetRegionName()+"%")
	}
	if req.GetCountryName() != "" {
		query = query.Where("country_name like ?", "%"+req.GetCountryName()+"%")
	}
	if req.GetAreaDesc() != "" {
		query = query.Where("area_desc like ?", "%"+req.GetAreaDesc()+"%")
	}
	if req.GetRegionDesc() != "" {
		query = query.Where("region_desc like ?", "%"+req.GetRegionDesc()+"%")
	}
	if req.GetCountryDesc() != "" {
		query = query.Where("country_desc like ?", "%"+req.GetCountryDesc()+"%")
	}
	return query
}

func QueryAreaList(ctx context.Context, req *system.SystemAreaRequest) (resp []*modle.SystemArea, err error) {
	query := getAreaQuery(ctx, req)
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
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryAreaById(ctx context.Context, id int64) (resp *modle.SystemArea, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(resp).First(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func QueryAreaByIdList(ctx context.Context, idList []int64) (resp []*modle.SystemArea, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemArea{}).Where("id in (?)", idList).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func CreateArea(ctx context.Context, entity *modle.SystemArea) error {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error
}

func BatchCreateArea(ctx context.Context, entityList []*modle.SystemArea) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(&modle.SystemArea{}).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateArea(ctx context.Context, entity *modle.SystemArea) error {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Save(entity).Error
}

func BatchUpdateArea(ctx context.Context, entityList []*modle.SystemArea) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, item := range entityList {
			if err = tx.Model(&modle.SystemArea{}).Where("id = ?", item.ID).Save(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
