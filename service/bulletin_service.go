package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryBulletinPage(ctx context.Context, req *system.SystemBulletinRequest) (resp *system.SystemBulletinPageResponse, err error) {
	list, total, err := repository.QueryBulletinPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*system.SystemBulletinResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &system.SystemBulletinPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp *system.SystemBulletinResponse, err error) {
	item, err := repository.QueryBulletinById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), nil
}

func SaveBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate().Int64()
	item := &modle.SystemBulletin{
		ID:        id,
		Title:     req.GetTitle(),
		Type:      req.Type,
		Status:    req.Status,
		Content:   req.GetContent(),
		IsDelete:  req.IsDelete,
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateBulletin(ctx, item); err != nil {
		return
	}
	return true, nil
}

func UpdateBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	item, err := repository.QueryBulletinById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.Title = req.GetTitle()
	item.Type = req.Type
	item.Content = req.GetContent()
	item.Status = req.Status
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateBulletin(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeleteBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	itemList, err := repository.QueryBulletinByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return
	}
	isDelete := int8(1)
	for _, bulletin := range itemList {
		bulletin.UpdatedBy = ctxutil.GetUserId(ctx)
		bulletin.IsDelete = &isDelete
	}
	if err = repository.BatchUpdateBulletin(ctx, itemList); err != nil {
		return
	}
	return true, nil
}
