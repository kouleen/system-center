package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryInterfacePage(ctx context.Context, req *system.SystemInterfaceRequest) (resp *system.SystemInterfacePageResponse, err error) {
	return service.QueryInterfacePage(ctx, req)
}

func QueryInterfaceList(ctx context.Context, req *system.SystemInterfaceRequest) (resp []*system.SystemInterfaceResponse, err error) {
	return service.QueryInterfaceList(ctx, req)
}

func QueryInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp *system.SystemInterfaceResponse, err error) {
	if req.Id == nil {
		return nil, errors.New("id is empty")
	}
	return service.QueryInterface(ctx, req)
}

func SaveInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if err = checkInterfaceCreate(req); err != nil {
		return false, err
	}
	return service.CreateInterface(ctx, req)
}

func UpdateInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if err = checkInterfaceUpdate(req); err != nil {
		return false, err
	}
	return service.UpdateInterface(ctx, req)
}

func DeleteInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if req.GetIdList() == nil || len(req.GetIdList()) == 0 {
		return false, errors.New("id is empty")
	}
	return service.DeleteInterface(ctx, req)
}

func CleanCacheInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	return service.CleanCacheInterface(ctx, req)
}

func checkInterfaceCreate(req *system.SystemInterfaceRequest) error {
	if req.RequestPath == "" {
		return errors.New("request path is empty")
	}
	if req.InterfaceName == "" {
		return errors.New("interface name is empty")
	}
	if req.MethodType == nil {
		return errors.New("method type is empty")
	}
	if req.MethodName == "" {
		return errors.New("method name is empty")
	}
	if req.Version == "" {
		return errors.New("version is empty")
	}
	if req.Status == nil {
		return errors.New("status is empty")
	}
	return nil
}

func checkInterfaceUpdate(req *system.SystemInterfaceRequest) error {
	if req.Id == nil {
		return errors.New("id is empty")
	}
	return checkInterfaceCreate(req)
}
