package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryLoginLogPage(ctx context.Context, req *system.SystemLoginLogRequest) (resp *system.SystemLoginLogPageResponse, err error) {
	return service.QueryLoginLogPage(ctx, req)
}

func ForcedRetreatLoginLog(ctx context.Context, req *system.SystemLoginLogRequest) (resp bool, err error) {
	if len(req.GetIdList()) == 0 {
		return false, errors.New("id list is empty")
	}
	return service.ForcedRetreatLoginLog(ctx, req)
}

func DeleteLoginLog(ctx context.Context, req *system.SystemLoginLogRequest) (resp bool, err error) {
	if len(req.GetIdList()) == 0 {
		return false, errors.New("id list is empty")
	}
	return service.DeleteLoginLog(ctx, req)
}
