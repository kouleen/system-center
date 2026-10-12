package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/client"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/idl/kitex_gen/user"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryOperateLogPage(ctx context.Context, req *system.SystemOperateLogRequest) (resp *system.SystemOperateLogPageResponse, err error) {
	list, total, err := repository.QueryOperateLogPage(ctx, req)
	if err != nil {
		return
	}
	records := make([]*system.SystemOperateLogResponse, len(list))
	userIdList := make([]int64, 0, len(list))
	for index, lis := range list {
		operateLogResponse := lis.ConvertResp()
		records[index] = operateLogResponse
		if operateLogResponse.CreatedBy > 0 {
			userIdList = append(userIdList, operateLogResponse.CreatedBy)
		}
	}
	userHeaderMap := map[int64]*user.UserHeaderResponse{}
	if len(userIdList) > 0 {
		rpcReq := &user.UserHeaderRequest{
			UserIdList: userIdList,
		}
		headerResponses, err := client.GetUserRpc().QueryUserHeaderList(ctx, rpcReq)
		if err == nil {
			for _, headerResponse := range headerResponses {
				userHeaderMap[headerResponse.Id] = headerResponse
			}
		}
	}
	for _, record := range records {
		if userHeader, ok := userHeaderMap[record.CreatedBy]; ok {
			record.Username = userHeader.Username
		}
	}
	return &system.SystemOperateLogPageResponse{
		Total:   total,
		Records: records,
	}, nil
}

func CreateOperateLog(ctx context.Context, req *system.SystemOperateLogRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate().Int64()
	item := &modle.SystemOperateLog{
		ID:        id,
		Module:    req.GetModule(),
		Business:  req.GetBusiness(),
		Type:      req.GetType(),
		IP:        req.GetIp(),
		Location:  req.GetLocation(),
		OS:        req.GetOs(),
		Browser:   req.GetBrowser(),
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedBy: req.GetCreatedBy(),
	}
	if err = repository.CreateOperateLog(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeleteOperateLog(ctx context.Context, req *system.SystemOperateLogRequest) (r bool, err error) {
	if err = repository.BatchDeleteOperateLog(ctx, req.GetIdList()); err != nil {
		return
	}
	return true, nil
}
