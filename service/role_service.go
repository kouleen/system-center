package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryRolePage(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRolePageResponse, err error) {
	list, total, err := repository.QueryRolePage(ctx, req)
	if err != nil {
		return
	}
	respList := make([]*system.SystemRoleResponse, len(list))
	for index, item := range list {
		respList[index] = item.ConvertResp()
	}
	return &system.SystemRolePageResponse{
		Total:   total,
		Records: respList,
	}, err
}

func QueryRole(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRoleResponse, err error) {
	item, err := repository.QueryRoleById(ctx, req.GetId())
	if err != nil {
		return
	}
	systemRoleMenus, err := repository.QueryRoleMenuByRoleId(ctx, req.GetId())
	if err != nil {
		return
	}
	itemResp := item.ConvertResp()
	if systemRoleMenus == nil || len(systemRoleMenus) == 0 {
		return itemResp, nil
	}
	menuIds := make([]int64, len(systemRoleMenus))

	for index, menu := range systemRoleMenus {
		menuIds[index] = menu.MenuId
	}
	itemResp.MenuIds = menuIds
	return itemResp, nil
}

func CreateRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	menuIdList := req.GetMenuIds()
	systemRoleMenus := make([]*modle.SystemRoleMenu, 0, len(menuIdList))
	if len(menuIdList) > 0 {
		for _, menuId := range menuIdList {
			systemRoleMenu := &modle.SystemRoleMenu{
				ID:        node.Generate().Int64(),
				MenuId:    menuId,
				RoleId:    id,
				CreatedBy: ctxutil.GetUserId(ctx),
				UpdatedBy: ctxutil.GetUserId(ctx),
			}
			systemRoleMenus = append(systemRoleMenus, systemRoleMenu)
		}
	}
	entity := &modle.SystemRole{
		ID:        id,
		RoleName:  req.GetRoleName(),
		RoleSort:  req.RoleSort,
		Status:    req.Status,
		Remark:    "",
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateRoleMenu(ctx, entity, systemRoleMenus); err != nil {
		return
	}

	return true, nil
}

func UpdateRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	item, err := repository.QueryRoleById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.RoleName = req.GetRoleName()
	item.RoleSort = req.RoleSort
	item.Status = req.Status
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	menuIdList := req.GetMenuIds()
	systemRoleMenus := make([]*modle.SystemRoleMenu, 0, len(menuIdList))
	if len(menuIdList) > 0 {
		for _, menuId := range menuIdList {
			systemRoleMenu := &modle.SystemRoleMenu{
				ID:        node.Generate().Int64(),
				MenuId:    menuId,
				RoleId:    item.ID,
				CreatedBy: ctxutil.GetUserId(ctx),
				UpdatedBy: ctxutil.GetUserId(ctx),
			}
			systemRoleMenus = append(systemRoleMenus, systemRoleMenu)
		}
	}
	if err = repository.UpdateRoleMenu(ctx, item, systemRoleMenus); err != nil {
		return
	}
	return true, nil
}
