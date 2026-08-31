package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/common/pkg/redis"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryTemplatePage(ctx context.Context, req *system.SystemTemplateRequest) (*system.SystemTemplatePageResponse, error) {
	list, total, err := repository.QueryTemplatePage(ctx, req)
	if err != nil {
		return nil, err
	}
	records := make([]*system.SystemTemplateResponse, len(list))
	for i, item := range list {
		var updateTime int64
		if item.UpdateTime != nil {
			updateTime = item.UpdateTime.UnixMilli()
		}
		records[i] = &system.SystemTemplateResponse{
			Id:              item.ID,
			TemplateCode:    item.TemplateCode,
			TemplateName:    item.TemplateName,
			TemplateType:    item.TemplateType,
			TemplateContent: item.TemplateContent,
			Remark:          item.Remark,
			IsDelete:        int8(item.IsDelete),
			CreatedBy:       item.CreatedBy,
			UpdatedBy:       item.UpdatedBy,
			CreateTime:      item.CreateTime.UnixMilli(),
			UpdateTime:      updateTime,
		}
	}
	return &system.SystemTemplatePageResponse{
		Total:   total,
		Records: records,
	}, nil
}

func CreateTemplate(ctx context.Context, req *system.SystemTemplateRequest) (bool, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	process := &redis.CodeProcess{}
	template := &modle.SystemTemplate{
		ID:              id,
		TemplateCode:    process.GenerateCode(ctx, &modle.TemplateCodeRule{}),
		TemplateName:    req.TemplateName,
		TemplateType:    req.TemplateType,
		TemplateContent: req.TemplateContent,
		Remark:          req.Remark,
		CreatedBy:       ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateTemplate(ctx, template); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateTemplate(ctx context.Context, req *system.SystemTemplateRequest) (resp bool, err error) {
	template, err := repository.QueryTemplateById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	template.TemplateName = req.TemplateName
	template.TemplateType = req.TemplateType
	template.TemplateContent = req.TemplateContent
	template.Remark = req.Remark
	template.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateTemplate(ctx, template); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteTemplate(ctx context.Context, req *system.SystemTemplateRequest) (resp bool, err error) {
	template, err := repository.QueryTemplateById(ctx, *req.Id)
	if err != nil {
		return false, err
	}
	template.IsDelete = uint8(*req.IsDelete)
	template.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateTemplate(ctx, template); err != nil {
		return false, err
	}
	return true, nil
}
