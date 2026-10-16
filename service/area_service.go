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

func QueryAreaTree(ctx context.Context, req *system.SystemAreaRequest) (resp []*system.SystemAreaResponse, err error) {
	list, err := repository.QueryAreaList(ctx, req)
	if err != nil {
		return
	}
	return buildAreaTree(list), nil
}

func buildAreaTree(systemAreas []*modle.SystemArea) []*system.SystemAreaResponse {
	// key: parentCode, value: 子区域列表
	groupMap := make(map[string][]*modle.SystemArea)
	// 1. 一次性分组，只遍历1次
	for _, area := range systemAreas {
		groupMap[area.ParentCode] = append(groupMap[area.ParentCode], area)
	}

	var areaList []*system.SystemAreaResponse
	// 2. 取出根节点 parentCode == ""
	rootList := groupMap[""]
	for _, root := range rootList {
		tree := root.ConvertResp()
		tree.Children = buildChildTree(root.AreaCode, groupMap)
		areaList = append(areaList, tree)
	}
	return areaList
}

// 递归从map拿子节点，不再循环全量数组
func buildChildTree(parentCode string, groupMap map[string][]*modle.SystemArea) []*system.SystemAreaResponse {
	childAreas, ok := groupMap[parentCode]
	if !ok {
		return nil
	}
	var children []*system.SystemAreaResponse
	for _, area := range childAreas {
		resp := area.ConvertResp()
		resp.Children = buildChildTree(area.AreaCode, groupMap)
		children = append(children, resp)
	}
	return children
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
