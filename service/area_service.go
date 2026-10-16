package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryAreaPage(ctx context.Context, req *system.SystemAreaRequest) (resp *system.SystemAreaPageResponse, err error) {
	list, total, err := repository.QueryAreaPage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*system.SystemAreaResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &system.SystemAreaPageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryAreaList(ctx context.Context, req *system.SystemAreaRequest) (resp []*system.SystemAreaResponse, err error) {
	list, err := repository.QueryAreaList(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*system.SystemAreaResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return respList, err
}

func QueryArea(ctx context.Context, request *system.SystemAreaRequest) (res *system.SystemAreaResponse, err error) {
	resp, err := repository.QueryAreaById(ctx, request.GetId())
	if err != nil {
		return
	}
	return resp.ConvertResp(), err
}

func SaveArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	entity := &modle.SystemArea{
		ID:            id,
		AreaCode:      request.GetAreaCode(),
		AreaName:      request.GetAreaName(),
		AreaDesc:      request.GetAreaDesc(),
		ParentCode:    request.GetParentCode(),
		ParentName:    request.GetParentName(),
		RegionCode:    request.GetRegionCode(),
		RegionName:    request.GetRegionName(),
		RegionDesc:    request.GetRegionDesc(),
		CountryCode:   request.GetCountryCode(),
		CountryName:   request.GetCountryName(),
		CountryDesc:   request.GetCountryDesc(),
		PhoneCode:     request.GetPhoneCode(),
		AreaLevelCode: request.GetAreaLevelCode(),
		Remark:        request.GetRemark(),
		CreatedBy:     ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateArea(ctx, entity); err != nil {
		return false, err
	}
	return true, nil
}

func UpdateArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	entity, err := repository.QueryAreaById(ctx, request.GetId())
	if err != nil {
		return false, err
	}
	entity.AreaCode = request.GetAreaCode()
	entity.AreaName = request.GetAreaName()
	entity.AreaDesc = request.GetAreaDesc()
	entity.ParentCode = request.GetParentCode()
	entity.ParentName = request.GetParentName()
	entity.RegionCode = request.GetRegionCode()
	entity.RegionName = request.GetRegionName()
	entity.RegionDesc = request.GetRegionDesc()
	entity.CountryCode = request.GetCountryCode()
	entity.CountryName = request.GetCountryName()
	entity.CountryDesc = request.GetCountryDesc()
	entity.PhoneCode = request.GetPhoneCode()
	entity.AreaLevelCode = request.GetAreaLevelCode()
	entity.Remark = request.GetRemark()
	entity.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateArea(ctx, entity); err != nil {
		return false, err
	}
	return true, nil
}

func DeleteArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	entityList, err := repository.QueryAreaByIdList(ctx, request.GetIdList())
	if err != nil || len(request.GetIdList()) != len(entityList) {
		return false, err
	}
	isDelete := int8(1)
	for _, item := range entityList {
		item.IsDelete = &isDelete
		item.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.BatchUpdateArea(ctx, entityList); err != nil {
		return false, err
	}
	return true, nil
}
