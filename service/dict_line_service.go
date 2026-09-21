package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryDictLinePage(ctx context.Context, req *system.SystemDictLineRequest) (resp *system.SystemDictLinePageResponse, err error) {
	list, total, err := repository.QueryDictLinePage(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemDictLineResponse, len(list))
	for index, dictLine := range list {
		records[index] = dictLine.ConvertResp()
	}
	return &system.SystemDictLinePageResponse{
		Total:   total,
		Records: records,
	}, nil
}

func QueryDictLineList(ctx context.Context, req *system.SystemDictLineRequest) (resp []*system.SystemDictLineResponse, err error) {
	list, err := repository.QueryDictLineList(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemDictLineResponse, len(list))
	for index, dictLine := range list {
		records[index] = dictLine.ConvertResp()
	}
	return records, nil
}

func QueryDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp *system.SystemDictLineResponse, err error) {
	systemDictLine, err := repository.QueryDictLineById(ctx, systemDictLineRequest.GetId())
	if err != nil {
		return
	}
	return systemDictLine.ConvertResp(), nil
}

func CreateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	line := &modle.SystemDictLine{
		ID:        id,
		DictCode:  req.GetDictCode(),
		DictValue: req.GetDictValue(),
		DictSort:  req.DictSort,
		DictType:  req.GetDictType(),
		ListClass: req.GetListClass(),
		Status:    req.Status,
		Remark:    req.GetRemark(),
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateDictLine(ctx, line); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	line, err := repository.QueryDictLineById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	line.DictCode = req.DictCode
	line.DictValue = req.DictValue
	line.DictSort = req.DictSort
	line.DictType = req.DictType
	line.ListClass = req.ListClass
	line.Status = req.Status
	line.Remark = req.Remark
	line.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateDictLine(ctx, line); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteDictLine(ctx context.Context, req *system.SystemDictLineRequest) (resp bool, err error) {
	line, err := repository.QueryDictLineById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	line.IsDelete = req.IsDelete
	line.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateDictLine(ctx, line); err != nil {
		return false, err
	}
	return true, nil
}
