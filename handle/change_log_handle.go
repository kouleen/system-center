package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryChangeLogPage(ctx context.Context, systemChangeLogRequest *system.SystemChangeLogRequest) (r *system.SystemChangeLogPageResponse, err error) {
	return service.QueryChangeLogPage(ctx, systemChangeLogRequest)
}

func DeleteChangeLog(ctx context.Context, systemChangeLogRequest *system.SystemChangeLogRequest) (r bool, err error) {
	if len(systemChangeLogRequest.GetIdList()) == 0 {
		return false, errors.New("id list is empty")
	}
	return service.DeleteChangeLog(ctx, systemChangeLogRequest)
}
