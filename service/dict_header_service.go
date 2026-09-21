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
		systemDictHeaderList[index] = dictHeader.ConvertResp()
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
		systemDictHeaderList[index] = dictHeader.ConvertResp()
	}
	return systemDictHeaderList, nil
}

func QueryDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderResponse, err error) {
	dictHeader, err := repository.QueryDictHeaderById(ctx, systemDictHeaderRequest.GetId())
	if err != nil {
		return
	}
	return dictHeader.ConvertResp(), err
}

func CreateDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	header := &modle.SystemDictHeader{
		ID:        id,
		DictName:  req.GetDictName(),
		DictType:  req.GetDictType(),
		Status:    req.Status,
		CreatedBy: ctxutil.GetUserId(ctx),
		Remark:    req.GetRemark(),
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
	header.Status = req.Status
	header.Remark = req.Remark
	header.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateDictHeader(ctx, header); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteDictHeader(ctx context.Context, req *system.SystemDictHeaderRequest) (resp bool, err error) {
	dictHeaderList, err := repository.QueryDictHeaderByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(dictHeaderList) {
		return false, err
	}
	isDelete := int8(1)
	for _, header := range dictHeaderList {
		header.IsDelete = &isDelete
		header.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateDictHeader(ctx, dictHeaderList); err != nil {
		return false, err
	}
	return true, nil
}
