package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryInterfacePage(ctx context.Context, req *system.SystemInterfaceRequest) (resp *system.SystemInterfacePageResponse, err error) {
	list, total, err := repository.QueryInterfacePage(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemInterfaceResponse, len(list))
	for index, lis := range list {
		records[index] = lis.ConvertResp()
	}
	return &system.SystemInterfacePageResponse{
		Total:   total,
		Records: records,
	}, nil
}

func QueryInterfaceList(ctx context.Context, req *system.SystemInterfaceRequest) (resp []*system.SystemInterfaceResponse, err error) {
	list, err := repository.QueryInterfaceList(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemInterfaceResponse, len(list))
	for index, lis := range list {
		records[index] = lis.ConvertResp()
	}
	return records, nil
}

func QueryInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp *system.SystemInterfaceResponse, err error) {
	systemInterface, err := repository.QueryInterfaceById(ctx, req.GetId())
	if err != nil {
		return
	}
	return systemInterface.ConvertResp(), nil
}

func CreateInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	line := &modle.SystemInterface{
		ID:            id,
		RequestPath:   req.GetRequestPath(),
		InterfaceName: req.GetInterfaceName(),
		MethodType:    req.MethodType,
		MethodName:    req.GetMethodName(),
		ParamTypes:    req.GetParamTypes(),
		Version:       req.GetVersion(),
		Status:        req.Status,
		Remark:        req.GetRemark(),
		CreatedBy:     ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateInterface(ctx, line); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	inter, err := repository.QueryInterfaceById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	inter.RequestPath = req.RequestPath
	inter.InterfaceName = req.InterfaceName
	inter.MethodType = req.MethodType
	inter.MethodName = req.MethodName
	inter.ParamTypes = req.ParamTypes
	inter.Status = req.Status
	inter.Remark = req.Remark
	inter.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateInterface(ctx, inter); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	inter, err := repository.QueryInterfaceById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	inter.IsDelete = req.IsDelete
	inter.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateInterface(ctx, inter); err != nil {
		return false, err
	}
	return true, nil
}

func CleanCacheInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	return repository.CleanCacheInterface(ctx, req)
}
