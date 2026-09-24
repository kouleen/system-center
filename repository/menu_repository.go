package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"gorm.io/gorm"
)

func QueryMenuListByUserId(ctx context.Context, userId int64) (resp []*modle.SystemMenu, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Table("system_menu t1").
		Joins("INNER JOIN system_role_menu t2 ON t1.id = t2.menu_id").
		Joins("INNER JOIN system_role_user t3 ON t2.role_id = t3.role_id").
		Select("t1.*").Where("t1.is_delete = 0 AND t3.user_id = ? order by order_num", userId).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryMenuList(ctx context.Context, req *system.SystemMenuRequest) (resp []*modle.SystemMenu, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Where("is_delete = ?", 0)
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	} else {
		query = query.Where("status = ?", 1)
	}
	if req.MenuName != "" {
		query = query.Where("menu_name like ?", "%"+req.MenuName+"%")
	}
	query = query.Order("order_num")
	if err = query.Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryMenuById(ctx context.Context, id int64) (resp *modle.SystemMenu, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Where("id = ?", id).Find(&resp).Error; err != nil {
		return
	}
	return
}

func QueryMenuByIdList(ctx context.Context, ids []int64) (resp []*modle.SystemMenu, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Where("id in (?)", ids).Find(&resp).Error; err != nil {
		return
	}
	return
}

func CreateMenu(ctx context.Context, entity *modle.SystemMenu) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Create(entity).Error
}

func BatchCreateMenu(ctx context.Context, entityList []*modle.SystemMenu) (err error) {
	return mysql.GetReadMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if len(entityList) > 0 {
			if err = tx.Model(&modle.SystemMenu{}).Create(entityList).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func UpdateMenu(ctx context.Context, entity *modle.SystemMenu) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Model(&modle.SystemMenu{}).Where("id = ?", entity.ID).Save(entity).Error
}

func BatchUpdateMenu(ctx context.Context, entityList []*modle.SystemMenu) (err error) {
	return mysql.GetWriteMysqlDDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, menu := range entityList {
			if err = tx.Model(&modle.SystemMenu{}).Where("id = ?", menu.ID).Save(menu).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
