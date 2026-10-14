package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/utils"
	"gorm.io/gorm"
)

func QueryRolePage(ctx context.Context, req *system.SystemRoleRequest) (resp []*modle.SystemRole, total int64, err error) {
	query := getRoleQuery(ctx, req)
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

func getRoleQuery(ctx context.Context, req *system.SystemRoleRequest) *gorm.DB {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemRole)).Where("is_delete = 0")
	if req.Status != nil {
		query.Where("status = ?", req.GetStatus())
	}
	if req.GetRoleName() != "" {
		query.Where("role_name LIKE ?", "%"+req.GetRoleName()+"%")
	}
	return query
}

func QueryRoleList(ctx context.Context, req *system.SystemRoleRequest) (resp []*modle.SystemRole, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemRole)).Table("system_role t1").
		Joins("INNER JOIN system_role_user t2 ON t2.role_id = t1.id").
		Select("t1.*").Where("t1.is_delete = 0 AND t2.user_id = ? order by role_sort", req.GetUserId()).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func QueryRoleById(ctx context.Context, id int64) (resp *modle.SystemRole, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&resp, id).Error; err != nil {
		return nil, err
	}
	return
}

func QueryRoleByIdList(ctx context.Context, ids []int64) (resp []*modle.SystemRole, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(new(modle.SystemRole)).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateRole(ctx context.Context, entity *modle.SystemRole) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(entity).Where("id = ?", entity.ID).Save(entity).Error
}
