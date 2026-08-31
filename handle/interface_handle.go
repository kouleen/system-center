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

func SaveInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if req.RequestPath == "" {
		return false, errors.New("request path is empty")
	}
	if req.InterfaceName == "" {
		return false, errors.New("interface name is empty")
	}
	if req.MethodType == nil {
		return false, errors.New("method type is empty")
	}
	if req.MethodName == "" {
		return false, errors.New("method name is empty")
	}
	if req.Version == "" {
		return false, errors.New("version is empty")
	}
	if req.Status == nil {
		return false, errors.New("status is empty")
	}
	return service.CreateInterface(ctx, req)
}

func UpdateInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is empty")
	}
	if req.RequestPath == "" {
		return false, errors.New("request path is empty")
	}
	if req.InterfaceName == "" {
		return false, errors.New("interface name is empty")
	}
	if req.MethodType == nil {
		return false, errors.New("method type is empty")
	}
	if req.MethodName == "" {
		return false, errors.New("method name is empty")
	}
	if req.Version == "" {
		return false, errors.New("version is empty")
	}
	if req.Status == nil {
		return false, errors.New("status is empty")
	}
	return service.UpdateInterface(ctx, req)
}

func DeleteInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is empty")
	}
	if req.IsDelete == nil {
		return false, errors.New("isDelete is empty")
	}
	return service.DeleteInterface(ctx, req)
}

func CleanCacheInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	return service.CleanCacheInterface(ctx, req)
}
