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
		var updateTime int64
		if lis.UpdateTime != nil {
			updateTime = lis.UpdateTime.UnixMilli()
		}
		records[index] = &system.SystemInterfaceResponse{
			Id:            lis.ID,
			RequestPath:   lis.RequestPath,
			InterfaceName: lis.InterfaceName,
			MethodType:    int8(lis.MethodType),
			MethodName:    lis.MethodName,
			ParamTypes:    lis.ParamTypes,
			Version:       lis.Version,
			Status:        int8(lis.Status),
			Remark:        lis.Remark,
			IsDelete:      int8(lis.IsDelete),
			CreatedBy:     lis.CreatedBy,
			UpdatedBy:     lis.UpdatedBy,
			CreateTime:    lis.CreateTime.UnixMilli(),
			UpdateTime:    updateTime,
		}
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
		var updateTime int64
		if lis.UpdateTime != nil {
			updateTime = lis.UpdateTime.UnixMilli()
		}
		records[index] = &system.SystemInterfaceResponse{
			Id:            lis.ID,
			RequestPath:   lis.RequestPath,
			InterfaceName: lis.InterfaceName,
			MethodType:    int8(lis.MethodType),
			MethodName:    lis.MethodName,
			ParamTypes:    lis.ParamTypes,
			Version:       lis.Version,
			Status:        int8(lis.Status),
			Remark:        lis.Remark,
			IsDelete:      int8(lis.IsDelete),
			CreatedBy:     lis.CreatedBy,
			UpdatedBy:     lis.UpdatedBy,
			CreateTime:    lis.CreateTime.UnixMilli(),
			UpdateTime:    updateTime,
		}
	}
	return records, nil
}

func CreateInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	line := &modle.SystemInterface{
		ID:            id,
		RequestPath:   req.RequestPath,
		InterfaceName: req.InterfaceName,
		MethodType:    uint8(*req.MethodType),
		MethodName:    req.MethodName,
		ParamTypes:    req.ParamTypes,
		Version:       req.Version,
		Status:        uint8(*req.Status),
		Remark:        req.Remark,
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
	inter.MethodType = uint8(*req.MethodType)
	inter.MethodName = req.MethodName
	inter.ParamTypes = req.ParamTypes
	inter.Status = uint8(*req.Status)
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
	inter.IsDelete = uint8(*req.IsDelete)
	inter.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateInterface(ctx, inter); err != nil {
		return false, err
	}
	return true, nil
}

func CleanCacheInterface(ctx context.Context, req *system.SystemInterfaceRequest) (resp bool, err error) {
	return repository.CleanCacheInterface(ctx, req)
}
