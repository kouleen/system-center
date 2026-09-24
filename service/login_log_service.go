package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryLoginLogPage(ctx context.Context, req *system.SystemLoginLogRequest) (resp *system.SystemLoginLogPageResponse, err error) {
	list, total, err := repository.QueryLoginLogPage(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemLoginLogResponse, len(list))
	for index, lis := range list {
		records[index] = lis.ConvertResp()
	}
	return &system.SystemLoginLogPageResponse{
		Total:   total,
		Records: records,
	}, nil
}

func CreateLoginLog(ctx context.Context, req *system.SystemLoginLogRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate().Int64()
	item := &modle.SystemLoginLog{
		ID:       id,
		Username: req.GetUsername(),
		IP:       req.GetIp(),
		Location: req.GetLocation(),
		OS:       req.GetOs(),
		Token:    req.GetToken(),
		Browser:  req.GetBrowser(),
		Status:   req.Status,
		Remark:   req.Remark,
	}
	if err = repository.CreateLoginLog(ctx, item); err != nil {
		return
	}
	return true, nil
}

func ForcedRetreatLoginLog(ctx context.Context, req *system.SystemLoginLogRequest) (resp bool, err error) {
	itemList, err := repository.QueryLoginLogByIdList(ctx, req.GetIdList())
	if err != nil || len(itemList) == 0 {
		return
	}
	tokenList := make([]string, len(itemList))
	for index, item := range itemList {
		tokenList[index] = item.Token
	}
	count, err := redis.Unlink(ctx, tokenList...)
	if err != nil {
		return
	}
	return count == int64(len(tokenList)), nil
}

func DeleteLoginLog(ctx context.Context, req *system.SystemLoginLogRequest) (resp bool, err error) {
	if err = repository.BatchDeleteLoginLog(ctx, req.GetIdList()); err != nil {
		return
	}
	return true, nil
}
