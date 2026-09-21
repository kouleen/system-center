package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryDictLinePage(ctx context.Context, req *system.SystemDictLineRequest) (resp *system.SystemDictLinePageResponse, err error) {
	return service.QueryDictLinePage(ctx, req)
}

func QueryDictLineList(ctx context.Context, req *system.SystemDictLineRequest) (resp []*system.SystemDictLineResponse, err error) {
	return service.QueryDictLineList(ctx, req)
}

func QueryDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp *system.SystemDictLineResponse, err error) {
	return service.QueryDictLine(ctx, systemDictLineRequest)
}

func CreateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if err = checkDictLineCreate(req); err != nil {
		return false, err
	}
	return service.CreateDictLine(ctx, req)
}

func UpdateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if err = checkDictLineUpdate(req); err != nil {
		return false, err
	}
	return service.UpdateDictLine(ctx, req)
}

func DeleteDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if req.GetIdList() == nil || len(req.GetIdList()) == 0 {
		return false, errors.New("id is required")
	}
	return service.DeleteDictLine(ctx, req)
}

func checkDictLineUpdate(req *system.SystemDictLineRequest) error {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkDictLineCreate(req)
}

func checkDictLineCreate(req *system.SystemDictLineRequest) error {
	if req.DictType == "" {
		return errors.New("dictType is required")
	}
	if req.DictCode == "" {
		return errors.New("dictCode is required")
	}
	if req.DictValue == "" {
		return errors.New("dictValue is required")
	}
	if req.DictSort == nil {
		return errors.New("dictSort is required")
	}
	if req.Status == nil {
		return errors.New("status is required")
	}
	return nil
}
