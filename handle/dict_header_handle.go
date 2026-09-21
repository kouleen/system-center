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

func QueryDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderResponse, err error) {
	return service.QueryDictHeader(ctx, systemDictHeaderRequest)
}

func CreateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if err = checkDictHeaderCreate(req); err != nil {
		return false, err
	}
	return service.CreateDictHeader(ctx, req)
}

func UpdateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if err = checkDictHeaderUpdate(req); err != nil {
		return false, err
	}
	return service.UpdateDictHeader(ctx, req)
}

func DeleteDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	if req.GetIdList() == nil || len(req.GetIdList()) == 0 {
		return false, errors.New("id is required")
	}
	return service.DeleteDictHeader(ctx, req)
}

func checkDictHeaderUpdate(req *system.SystemDictHeaderRequest) error {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkDictHeaderCreate(req)
}

func checkDictHeaderCreate(req *system.SystemDictHeaderRequest) error {
	if req.GetDictName() == "" {
		return errors.New("dict name is required")
	}
	if req.GetDictType() == "" {
		return errors.New("dict type is required")
	}
	if req.Status == nil {
		return errors.New("status is required")
	}
	return nil
}
