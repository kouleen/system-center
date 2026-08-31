package repository

import (
	"context"

	"github.com/kouleen/common/pkg/mysql"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
)

func QueryInterfacePage(ctx context.Context, req *system.SystemInterfaceRequest) (list []*modle.SystemInterface, total int64, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Where("is_delete = ?", 0)
	if req.RequestPath != "" {
		query = query.Where("request_path LIKE = ?", req.RequestPath+"%")
	}
	if req.InterfaceName != "" {
		query = query.Where("interface_name LIKE = ?", req.InterfaceName+"%")
	}
	if req.MethodType != nil {
		query = query.Where("method_type = ?", req.MethodType)
	}
	if req.MethodName != "" {
		query = query.Where("method_name LIKE = ?", req.MethodName+"%")
	}
	if req.Version != "" {
		query = query.Where("version = ?", req.Version)
	}
	if err = query.Count(&total).Error; err != nil {
		return
	}
	query = query.Order("create_time desc")
	i := (req.GetCurrent() - 1) * req.GetSize()
	if err = query.Offset(int(i)).Limit(int(req.GetSize())).Find(&list).Error; err != nil {
		return
	}
	return
}

func QueryInterfaceList(ctx context.Context, req *system.SystemInterfaceRequest) (list []*modle.SystemInterface, err error) {
	query := mysql.GetReadMysqlDDB().WithContext(ctx).Where("is_delete = ?", 0)
	if req.RequestPath != "" {
		query = query.Where("request_path LIKE ?", req.RequestPath+"%")
	}
	if req.InterfaceName != "" {
		query = query.Where("interface_name LIKE ?", req.InterfaceName+"%")
	}
	if req.MethodType != nil {
		query = query.Where("method_type = ?", req.MethodType)
	}
	if req.MethodName != "" {
		query = query.Where("method_name LIKE = ?", req.MethodName+"%")
	}
	if req.Version != "" {
		query = query.Where("version = ?", req.Version)
	}
	query = query.Order("create_time desc")
	if err = query.Find(&list).Error; err != nil {
		return
	}
	return
}

func QueryInterfaceById(ctx context.Context, id int64) (systemInterface *modle.SystemInterface, err error) {
	if err = mysql.GetReadMysqlDDB().WithContext(ctx).First(&systemInterface, id).Error; err != nil {
		return
	}
	return
}

func CreateInterface(ctx context.Context, entity *modle.SystemInterface) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Create(entity).Error; err != nil {
		return
	}
	return
}

func UpdateInterface(ctx context.Context, entity *modle.SystemInterface) (err error) {
	if err = mysql.GetWriteMysqlDDB().WithContext(ctx).Where("id = ?", entity.ID).Updates(entity).Error; err != nil {
		return
	}
	return
}

func CleanCacheInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	return
}
