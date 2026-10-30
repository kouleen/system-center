package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryChangeLogPage(ctx context.Context, req *system.SystemChangeLogRequest) (r *system.SystemChangeLogPageResponse, err error) {
	list, total, err := repository.QueryChangeLogPage(ctx, req)
	if err != nil {
		return
	}
	itemList := make([]*system.SystemChangeLogResponse, len(list))
	for index, item := range list {
		itemList[index] = item.ConvertResp()
	}
	return &system.SystemChangeLogPageResponse{
		Total:   total,
		Records: itemList,
	}, err
}

func BatchCreateChangeLog(ctx context.Context, reqList []*system.SystemChangeLogRequest) (err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	entityList := make([]*modle.SystemChangeLog, len(reqList))
	for index, req := range reqList {
		id := node.Generate().Int64()
		entityList[index] = &modle.SystemChangeLog{
			ID:     id,
			Schema: req.GetSchema(),
			Table:  req.GetTable(),
			Action: req.GetAction(),
			Before: req.GetBefore(),
			After:  req.GetAfter(),
			Remark: req.Remark,
		}
	}
	return repository.BatchCreateChangeLog(ctx, entityList)
}

func DeleteChangeLog(ctx context.Context, req *system.SystemChangeLogRequest) (r bool, err error) {
	if err = repository.BatchDeleteChangeLog(ctx, req.GetIdList()); err != nil {
		return
	}
	return true, nil
}
