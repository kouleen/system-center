package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/utils"
)

func UploadFile(ctx context.Context, req *system.SystemFileRequest) (resp *system.SystemFileResponse, err error) {
	return utils.UploadFile(ctx, req)
}
