package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryMenuTree(ctx context.Context, req *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	return service.QueryMenuTree(ctx, req)
}

func QueryMenuList(ctx context.Context, req *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	return service.QueryMenuList(ctx, req)
}

func QueryMenu(ctx context.Context, req *system.SystemMenuRequest) (resp *system.SystemMenuResponse, err error) {
	if req.Id == nil {
		return nil, errors.New("id is required")
	}
	return service.QueryMenu(ctx, req)
}

func SaveMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	if err = checkMenuCreate(req); err != nil {
		return false, err
	}
	return service.SaveMenu(ctx, req)
}

func UpdateMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	if err = checkMenuUpdate(req); err != nil {
		return false, err
	}
	return service.UpdateMenu(ctx, req)
}

func DeleteMenu(ctx context.Context, req *system.SystemMenuRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("id is required")
	}
	return service.DeleteMenu(ctx, req)
}

func checkMenuCreate(req *system.SystemMenuRequest) error {
	if req.ParentId == nil {
		return errors.New("id is required")
	}
	if req.MenuType == nil {
		return errors.New("menuType is required")
	}
	if req.GetMenuName() == "" {
		return errors.New("menuName is required")
	}
	if req.OrderNum == nil {
		return errors.New("orderNum is required")
	}
	return nil
}

func checkMenuUpdate(req *system.SystemMenuRequest) error {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkMenuCreate(req)
}
