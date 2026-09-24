package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryRoleMenuByRoleId(ctx context.Context, roleId int64) (resp []*modle.SystemRoleMenu, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Where("role_id = ?", roleId).Find(&resp).Error; err != nil {
		return
	}
	return
}

func CreateRoleMenu(ctx context.Context, systemRole *modle.SystemRole, systemRoleMenus []*modle.SystemRoleMenu) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&modle.SystemRole{}).Create(systemRole).Error; err != nil {
			return err
		}
		if len(systemRoleMenus) > 0 {
			if err = tx.Model(&modle.SystemRoleMenu{}).Create(systemRoleMenus).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateRoleMenu(ctx context.Context, systemRole *modle.SystemRole, systemRoleMenus []*modle.SystemRoleMenu) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&modle.SystemRole{}).Where("id = ?", systemRole.ID).Save(systemRole).Error; err != nil {
			return err
		}
		if err = tx.Where("role_id = ?", systemRole.ID).Delete(&modle.SystemRoleMenu{}).Error; err != nil {
			return err
		}
		if len(systemRoleMenus) > 0 {
			if err = tx.Model(&modle.SystemRoleMenu{}).Create(systemRoleMenus).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteRoleMenu(ctx context.Context, roleIdList []int64, systemRoleList []*modle.SystemRole) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(roleIdList) > 0 {
			if err = tx.Where("role_id in (?)", roleIdList).Delete(&modle.SystemRoleMenu{}).Error; err != nil {
				return err
			}
		}
		if len(systemRoleList) > 0 {
			for _, role := range systemRoleList {
				if err = tx.Model(&modle.SystemRole{}).Where("id = ?", role.ID).Save(role).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
