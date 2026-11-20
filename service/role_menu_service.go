package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryRoleMenuList(ctx context.Context, req *system.SystemRoleMenuRequest) (resp []*system.SystemRoleMenuResponse, err error) {
	list, err := repository.QueryRoleMenuByRoleId(ctx, req.GetRoleId())
	if err != nil {
		return
	}
	records := make([]*system.SystemRoleMenuResponse, len(list))
	for index, dictLine := range list {
		records[index] = dictLine.ConvertResp()
	}
	return records, nil
}

func SaveRoleMenu(ctx context.Context, req *system.SystemRoleMenuRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	menuIdList := req.GetMenuIds()
	systemRoleMenus := make([]*modle.SystemRoleMenu, 0, len(menuIdList))
	if len(menuIdList) > 0 {
		for _, menuId := range menuIdList {
			systemRoleMenu := &modle.SystemRoleMenu{
				ID:        node.Generate().Int64(),
				MenuId:    menuId,
				RoleId:    req.GetRoleId(),
				CreatedBy: ctxutil.GetUserId(ctx),
				UpdatedBy: ctxutil.GetUserId(ctx),
			}
			systemRoleMenus = append(systemRoleMenus, systemRoleMenu)
		}
	}
	if err = repository.SaveRoleMenu(ctx, req.GetRoleId(), systemRoleMenus); err != nil {
		return
	}
	return true, nil
}
