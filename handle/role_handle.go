package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
)

func QueryRolePage(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRolePageResponse, err error) {
	return
}

func QueryRole(ctx context.Context, req *system.SystemRoleRequest) (resp *system.SystemRoleResponse, err error) {
	return
}

func SaveRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	return
}

func UpdateRole(ctx context.Context, req *system.SystemRoleRequest) (resp bool, err error) {
	return
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
