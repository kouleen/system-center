package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryEducationPage(ctx context.Context, req *system.SystemEducationRequest) (resp *system.SystemEducationPageResponse, err error) {
	list, total, err := repository.QueryEducationPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*system.SystemEducationResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &system.SystemEducationPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryEducation(ctx context.Context, req *system.SystemEducationRequest) (resp *system.SystemEducationResponse, err error) {
	item, err := repository.QueryEducationById(ctx, req.GetId())
	if err != nil {
		return
	}
	return item.ConvertResp(), nil
}

func SaveEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	id := node.Generate().Int64()
	item := &modle.SystemEducation{
		ID:        id,
		Title:     req.GetTitle(),
		Hits:      req.Hits,
		Remark:    req.GetRemark(),
		IsDelete:  req.IsDelete,
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateEducation(ctx, item); err != nil {
		return
	}
	return true, nil
}

func UpdateEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	item, err := repository.QueryEducationById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.Title = req.GetTitle()
	item.Hits = req.Hits
	item.Remark = req.GetRemark()
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateEducation(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeleteEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	itemList, err := repository.QueryEducationByIdList(ctx, req.GetIdList())
	if err != nil || len(req.GetIdList()) != len(itemList) {
		return
	}
	isDelete := int8(1)
	for _, bulletin := range itemList {
		bulletin.UpdatedBy = ctxutil.GetUserId(ctx)
		bulletin.IsDelete = &isDelete
	}
	if err = repository.BatchCrateEducation(ctx, itemList); err != nil {
		return
	}
	return true, nil
}
