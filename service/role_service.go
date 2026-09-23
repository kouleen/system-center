package service

import (
	"context"
	"errors"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/client"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/idl/kitex_gen/user"
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

func UpdateRoleStatus(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	item, err := repository.QueryRoleById(ctx, req.GetId())
	if err != nil {
		return
	}
	item.Status = req.Status
	item.UpdatedBy = ctxutil.GetUserId(ctx)
	if err = repository.UpdateRole(ctx, item); err != nil {
		return
	}
	return true, nil
}

func DeleteRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	itemList, err := repository.QueryRoleByIdList(ctx, req.GetIdList())
	if err != nil {
		return
	}
	isDelete := int8(1)
	for _, role := range itemList {
		role.IsDelete = &isDelete
		role.UpdatedBy = ctxutil.GetUserId(ctx)
	}
	if err = repository.DeleteRoleMenu(ctx, req.GetIdList(), itemList); err != nil {
		return
	}
	return true, nil
}

func QueryRoleUserPage(ctx context.Context, req *system.SystemRoleUserRequest) (resp *system.SystemRoleUserPageResponse, err error) {
	roleUserList, err := repository.QueryRoleUserByRoleId(ctx, req.GetRoleId())
	if err != nil {
		return
	}
	if len(roleUserList) == 0 {
		return
	}
	userIdList := make([]int64, 0, len(roleUserList))
	for _, roleUser := range roleUserList {
		userIdList = append(userIdList, roleUser.UserId)
	}
	userHeaderReq := &user.UserHeaderRequest{
		Params:     req.GetParams(),
		Current:    req.GetCurrent(),
		Size:       req.GetSize(),
		Username:   req.GetUsername(),
		Nickname:   req.GetNickname(),
		Phone:      req.GetPhone(),
		UserIdList: userIdList,
	}
	userHeaderResp, err := client.GetUserRpc().QueryUserHeaderPage(ctx, userHeaderReq)
	if err != nil {
		return
	}
	records := userHeaderResp.Records
	respRecords := make([]*system.SystemRoleUserResponse, len(records))
	for index, record := range records {
		respRecords[index] = &system.SystemRoleUserResponse{
			Id:         record.Id,
			Username:   record.GetUsername(),
			Nickname:   record.GetNickname(),
			Gender:     record.Gender,
			Avatar:     record.GetAvatar(),
			Phone:      record.GetPhone(),
			Status:     record.Status,
			IsDelete:   record.IsDelete,
			CreatedBy:  record.CreatedBy,
			UpdatedBy:  record.UpdatedBy,
			CreateTime: record.CreateTime,
		}
	}
	return &system.SystemRoleUserPageResponse{
		Total:   userHeaderResp.Total,
		Records: respRecords,
	}, nil
}

func QueryRoleUserList(ctx context.Context, req *system.SystemRoleUserRequest) (resp []*system.SystemRoleUserResponse, err error) {
	itemList, err := repository.QueryRoleUserByRoleId(ctx, req.GetRoleId())
	if err != nil {
		return
	}
	if len(itemList) == 0 {
		return
	}
	respList := make([]*system.SystemRoleUserResponse, len(itemList))
	for index, item := range itemList {
		respList[index] = item.ConvertResp()
	}
	return respList, nil
}

func SaveRoleUser(ctx context.Context, req *system.SystemRoleUserRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return
	}
	role, err := repository.QueryRoleById(ctx, req.GetRoleId())
	if err != nil {
		return
	}
	if common.BaseStatus_DISABLED == common.BaseStatus(*role.Status) {
		return false, errors.New("role status is disabled")
	}

	roleUserList := make([]*modle.SystemRoleUser, len(req.GetUserIdList()))
	for index, userId := range req.GetUserIdList() {
		roleUserList[index] = &modle.SystemRoleUser{
			ID:        node.Generate().Int64(),
			RoleId:    role.ID,
			UserId:    userId,
			CreatedBy: ctxutil.GetUserId(ctx),
			UpdatedBy: ctxutil.GetUserId(ctx),
		}
	}
	if err = repository.BatchCreateRoleUser(ctx, req.GetRoleId(), roleUserList); err != nil {
		return
	}
	return true, nil
}

func CancelRoleUser(ctx context.Context, req *system.SystemRoleUserRequest) (resp bool, err error) {
	if err = repository.DeleteRoleUser(ctx, req.GetRoleId(), req.GetUserIdList()); err != nil {
		return
	}
	return true, nil
}
