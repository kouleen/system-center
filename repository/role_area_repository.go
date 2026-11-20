package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryRoleAreaByRoleId(ctx context.Context, roleId int64) (resp []*modle.SystemRoleArea, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Where("role_id = ?", roleId).Find(&resp).Error; err != nil {
		return nil, err
	}
	return
}

func SaveRoleArea(ctx context.Context, roleId int64, systemRoleAreas []*modle.SystemRoleArea) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err = tx.Where("role_id = ?", roleId).Delete(&modle.SystemRoleArea{}).Error; err != nil {
			return err
		}
		if len(systemRoleAreas) > 0 {
			if err = tx.Model(&modle.SystemRoleArea{}).Create(systemRoleAreas).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
