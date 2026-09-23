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
		if err = tx.Model(&modle.SystemRole{}).Where("id = ?", systemRole.ID).Updates(systemRole).Error; err != nil {
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
