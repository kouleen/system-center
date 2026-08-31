package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryDictHeaderPage(ctx context.Context, req *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderPageResponse, err error) {
	list, total, err := repository.QueryDictHeaderPage(ctx, req)
	if err != nil {
		return
	}
	systemDictHeaderList := make([]*system.SystemDictHeaderResponse, len(list))
	for index, dictHeader := range list {
		var updateTime int64
		if dictHeader.UpdateTime != nil {
			updateTime = dictHeader.UpdateTime.UnixMilli()
		}
		systemDictHeaderList[index] = &system.SystemDictHeaderResponse{
			Id:         dictHeader.ID,
			DictName:   dictHeader.DictName,
			DictType:   dictHeader.DictType,
			Status:     int8(dictHeader.Status),
			Remark:     dictHeader.Remark,
			IsDelete:   int8(dictHeader.IsDelete),
			CreatedBy:  dictHeader.CreatedBy,
			UpdatedBy:  dictHeader.UpdatedBy,
			CreateTime: dictHeader.CreateTime.UnixMilli(),
			UpdateTime: updateTime,
		}
	}
	return &system.SystemDictHeaderPageResponse{
		Total:   total,
		Records: systemDictHeaderList,
	}, err
}

func QueryDictHeaderList(ctx context.Context, req *system.SystemDictHeaderRequest) (resp []*system.SystemDictHeaderResponse, err error) {
	list, err := repository.QueryDictHeaderList(ctx, req)
	if err != nil {
		return
	}
	systemDictHeaderList := make([]*system.SystemDictHeaderResponse, len(list))
	for index, dictHeader := range list {
		var updateTime int64
		if dictHeader.UpdateTime != nil {
			updateTime = dictHeader.UpdateTime.UnixMilli()
		}
		header := &system.SystemDictHeaderResponse{
			Id:         dictHeader.ID,
			DictName:   dictHeader.DictName,
			DictType:   dictHeader.DictType,
			Status:     int8(dictHeader.Status),
			Remark:     dictHeader.Remark,
			IsDelete:   int8(dictHeader.IsDelete),
			CreatedBy:  dictHeader.CreatedBy,
			UpdatedBy:  dictHeader.UpdatedBy,
			CreateTime: dictHeader.CreateTime.UnixMilli(),
			UpdateTime: updateTime,
		}
		systemDictHeaderList[index] = header
	}
	return systemDictHeaderList, nil
}

func CreateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	header := &modle.SystemDictHeader{
		ID:        id,
		DictName:  req.DictName,
		DictType:  req.DictType,
		Status:    uint8(*req.Status),
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateDictHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	header, err := repository.QueryDictHeaderById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	header.DictName = req.DictName
	header.DictType = req.DictType
	header.Status = uint8(*req.Status)
	header.Remark = req.Remark
	header.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateDictHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	header, err := repository.QueryDictHeaderById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	header.IsDelete = uint8(*req.IsDelete)
	header.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateDictHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}
