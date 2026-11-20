package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryRoleMenuList(ctx context.Context, systemRoleMenuRequest *system.SystemRoleMenuRequest) (resp []*system.SystemRoleMenuResponse, err error) {
	return service.QueryRoleMenuList(ctx, systemRoleMenuRequest)
}

func SaveRoleMenu(ctx context.Context, req *system.SystemRoleMenuRequest) (resp bool, err error) {
	if req.RoleId == nil {
		return false, errors.New("role_id is required")
	}
	return service.SaveRoleMenu(ctx, req)
}
