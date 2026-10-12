package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryOperateLogPage(ctx context.Context, req *system.SystemOperateLogRequest) (resp *system.SystemOperateLogPageResponse, err error) {
	return service.QueryOperateLogPage(ctx, req)
}

func DeleteOperateLog(ctx context.Context, req *system.SystemOperateLogRequest) (r bool, err error) {
	if len(req.GetIdList()) == 0 {
		return false, errors.New("id list is empty")
	}
	return service.DeleteOperateLog(ctx, req)
}
