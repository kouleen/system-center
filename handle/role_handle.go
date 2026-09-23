package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryRolePage(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRolePageResponse, err error) {
	return service.QueryRolePage(ctx, req)
}

func QueryRole(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRoleResponse, err error) {
	return service.QueryRole(ctx, req)
}

func SaveRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	if err = checkCreateRole(req); err != nil {
		return
	}
	return service.CreateRole(ctx, req)
}

func UpdateRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	if err = checkUpdateRole(req); err != nil {
		return
	}
	return service.UpdateRole(ctx, req)
}

func DeleteRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	return
}

func SaveRoleMenu(ctx context.Context, req *system.SystemRoleMenuRequest) (resp bool, err error) {
	return
}

func UpdateRoleMenu(ctx context.Context, req *system.SystemRoleMenuRequest) (resp bool, err error) {
	return
}

func QuerySystemRoleUserPage(ctx context.Context, req *system.SystemRoleUserRequest) (resp *system.SystemRoleUserPageResponse, err error) {
	return
}

func QueryRoleUserList(ctx context.Context, req *system.SystemRoleUserRequest) (resp []*system.SystemRoleUserResponse, err error) {
	return
}

func SaveRoleUser(ctx context.Context, req *system.SystemRoleUserRequest) (resp bool, err error) {
	return
}

func CancelRoleUser(ctx context.Context, req *system.SystemRoleUserRequest) (resp bool, err error) {
	return
}

func checkUpdateRole(req *system.SystemRoleRequest) error {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkCreateRole(req)
}

func checkCreateRole(req *system.SystemRoleRequest) error {
	if req.GetRoleName() == "" {
		return errors.New("role_name is empty")
	}
	if req.RoleSort == nil {
		return errors.New("role_sort is empty")
	}
	if req.Status == nil {
		return errors.New("status is empty")
	}
	return nil
}
