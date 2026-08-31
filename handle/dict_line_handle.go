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

func CreateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if req.DictType == "" {
		return false, errors.New("dictType is required")
	}
	if req.DictCode == "" {
		return false, errors.New("dictCode is required")
	}
	if req.DictValue == "" {
		return false, errors.New("dictValue is required")
	}
	if req.DictSort == nil {
		return false, errors.New("dictSort is required")
	}
	if req.Status == nil {
		return false, errors.New("status is required")
	}
	return service.CreateDictLine(ctx, req)
}

func UpdateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is required")
	}
	if req.DictType == "" {
		return false, errors.New("dictType is required")
	}
	if req.DictCode == "" {
		return false, errors.New("dictCode is required")
	}
	if req.DictValue == "" {
		return false, errors.New("dictValue is required")
	}
	if req.DictSort == nil {
		return false, errors.New("dictSort is required")
	}
	if req.Status == nil {
		return false, errors.New("status is required")
	}
	return service.UpdateDictLine(ctx, req)
}

func DeleteDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is required")
	}
	if req.IsDelete == nil {
		return false, errors.New("isDelete is required")
	}
	return service.DeleteDictLine(ctx, req)
}
