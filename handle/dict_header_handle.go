package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryDictHeaderPage(ctx context.Context, req *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderPageResponse, err error) {
	return service.QueryDictHeaderPage(ctx, req)
}

func QueryDictHeaderList(ctx context.Context, req *system.SystemDictHeaderRequest) (resp []*system.SystemDictHeaderResponse, err error) {
	return service.QueryDictHeaderList(ctx, req)
}

func CreateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if req.DictName == "" {
		return false, errors.New("dict name is required")
	}
	if req.DictType == "" {
		return false, errors.New("dict type is required")
	}
	if req.Status == nil {
		return false, errors.New("status is required")
	}
	return service.CreateDictHeader(ctx, req)
}

func UpdateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if req.Id != nil {
		return false, errors.New("id is required")
	}
	if req.DictName == "" {
		return false, errors.New("dict name is required")
	}
	if req.DictType == "" {
		return false, errors.New("dict type is required")
	}
	if req.Status == nil {
		return false, errors.New("status is required")
	}
	return service.UpdateDictHeader(ctx, req)
}

func DeleteDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is required")
	}
	if req.IsDelete == nil {
		return false, errors.New("isDelete is required")
	}
	return service.DeleteDictHeader(ctx, req)
}
