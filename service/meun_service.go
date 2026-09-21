package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryMenuTree(ctx context.Context, req *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	systemMenus, err := repository.QueryMenuListByUserId(ctx, ctxutil.GetUserId(ctx))
	if err != nil {
		return
	}
	return buildMenuTree(systemMenus), nil
}

func QueryMenuList(ctx context.Context, req *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	systemMenus, err := repository.QueryMenuList(ctx, req)
	if err != nil {
		return
	}
	return buildMenuTree(systemMenus), nil
}

func QueryMenu(ctx context.Context, req *system.SystemMenuRequest) (resp *system.SystemMenuResponse, err error) {
	systemMenu, err := repository.QueryMenuById(ctx, req.GetId())
	if err != nil {
		return
	}
	return systemMenu.ConvertResp(), err
}

func SaveMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	id := node.Generate().Int64()
	entity := &modle.SystemMenu{
		ID:        id,
		MenuName:  req.GetMenuName(),
		ParentId:  req.GetParentId(),
		OrderNum:  req.OrderNum,
		Path:      req.GetPath(),
		Component: req.GetComponent(),
		Query:     req.GetQuery(),
		RouteName: req.GetRouteName(),
		IsFrame:   req.IsFrame,
		IsCache:   req.IsCache,
		MenuType:  req.MenuType,
		Visible:   req.Visible,
		Status:    req.Status,
		Perms:     req.GetPerms(),
		Icon:      req.GetIcon(),
		IsDelete:  req.IsDelete,
		CreatedBy: ctxutil.GetUserId(ctx),
	}
	if err = repository.CreateMenu(ctx, entity); err != nil {
		return
	}
	return true, nil
}

func UpdateMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	systemMenu, err := repository.QueryMenuById(ctx, req.GetId())
	if err != nil {
		return
	}
	systemMenu.MenuName = req.GetMenuName()
	systemMenu.ParentId = req.GetParentId()
	systemMenu.OrderNum = req.OrderNum
	systemMenu.Path = req.GetPath()
	systemMenu.Component = req.GetComponent()
	systemMenu.Query = req.GetQuery()
	systemMenu.RouteName = req.GetRouteName()
	systemMenu.IsFrame = req.IsFrame
	systemMenu.IsCache = req.IsCache
	systemMenu.MenuType = req.MenuType
	systemMenu.Visible = req.Visible
	systemMenu.Status = req.Status
	systemMenu.Perms = req.GetPerms()
	systemMenu.Icon = req.GetIcon()
	systemMenu.IsDelete = req.IsDelete
	systemMenu.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateMenu(ctx, systemMenu); err != nil {
		return
	}
	return true, nil
}

func DeleteMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	systemMenu, err := repository.QueryMenuById(ctx, req.GetId())
	if err != nil {
		return
	}
	isDelete := int8(1)
	systemMenu.IsDelete = &isDelete
	systemMenu.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateMenu(ctx, systemMenu); err != nil {
		return
	}
	return true, nil
}

func buildMenuTree(systemMenus []*modle.SystemMenu) []*system.SystemMenuResponse {
	var treeList []*system.SystemMenuResponse
	// 1. 先找出所有一级菜单（ParentId == 0）
	for _, menu := range systemMenus {
		if menu.ParentId == 0 {
			tree := menu.ConvertResp()
			// 2. 递归查找子菜单
			tree.Children = getChildren(menu.ID, systemMenus)
			treeList = append(treeList, tree)
		}
	}
	return treeList
}

func getChildren(parentId int64, menuList []*modle.SystemMenu) []*system.SystemMenuResponse {
	var children []*system.SystemMenuResponse
	for _, menu := range menuList {
		if menu.ParentId == parentId {
			tree := menu.ConvertResp()
			// 递归找子菜单
			tree.Children = getChildren(menu.ID, menuList)
			children = append(children, tree)
		}
	}
	return children
}
