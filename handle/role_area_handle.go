package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryRoleAreaList(ctx context.Context, systemRoleAreaRequest *system.SystemRoleAreaRequest) (resp []*system.SystemRoleAreaResponse, err error) {
	return service.QueryRoleAreaList(ctx, systemRoleAreaRequest)
}

func SaveRoleArea(ctx context.Context, req *system.SystemRoleAreaRequest) (resp bool, err error) {
	if req.RoleId == nil {
		return false, errors.New("role_id is required")
	}
	return service.SaveRoleArea(ctx, req)
}
