package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func UploadFile(ctx context.Context, systemFileRequest *system.SystemFileRequest) (resp *system.SystemFileResponse, err error) {
	return service.UploadFile(ctx, systemFileRequest)
}
