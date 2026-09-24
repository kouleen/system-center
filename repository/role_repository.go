package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
)

func QueryRolePage(ctx context.Context, req *system.SystemRoleRequest) (resp []*modle.SystemRole, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemRole)).Where("is_delete = 0")
	if req.Status != nil {
		query.Where("status = ?", req.GetStatus())
	}
	if req.GetRoleName() != "" {
		query.Where("role_name LIKE ?", "%"+req.GetRoleName()+"%")
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
	query = query.Order("role_sort")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryRoleById(ctx context.Context, id int64) (resp *modle.SystemRole, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&resp, id).Error; err != nil {
		return
	}
	return
}

func QueryRoleByIdList(ctx context.Context, ids []int64) (resp []*modle.SystemRole, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemRole)).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return
	}
	return
}

func UpdateRole(ctx context.Context, entity *modle.SystemRole) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(entity).Where("id = ?", entity.ID).Save(entity).Error
}
