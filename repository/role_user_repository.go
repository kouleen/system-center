package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryRoleUserByRoleId(ctx context.Context, roleId int64) (resp []*modle.SystemRoleUser, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemRoleUser{}).Where("role_id = ?", roleId).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryRoleUserByUserIdList(ctx context.Context, roleId int64, userIdList []int64) (resp []*modle.SystemRoleUser, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemRoleUser{}).Where("role_id = ? AND user_id in (?)", roleId, userIdList).Find(&resp).Error; err != nil {
		return
	}
	return
}

func BatchCreateRoleUser(ctx context.Context, roleId int64, roleUserList []*modle.SystemRoleUser) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("role_id = ?", roleId).Delete(&modle.SystemRoleUser{}).Error; err != nil {
			return err
		}
		if len(roleUserList) > 0 {
			if err = tx.Model(&modle.SystemRoleUser{}).Create(roleUserList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteRoleUser(ctx context.Context, roleId int64, userIdList []int64) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("role_id = ? and user_id in (?)", roleId, userIdList).Delete(&modle.SystemRoleUser{}).Error; err != nil {
			return err
		}
		return nil
	})
}
