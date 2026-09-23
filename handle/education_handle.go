package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryEducationPage(ctx context.Context, req *system.SystemEducationRequest) (resp *system.SystemEducationPageResponse, err error) {
	return service.QueryEducationPage(ctx, req)
}

func QueryEducation(ctx context.Context, req *system.SystemEducationRequest) (resp *system.SystemEducationResponse, err error) {
	return service.QueryEducation(ctx, req)
}

func SaveEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	return service.SaveEducation(ctx, req)
}

func UpdateEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	return service.UpdateEducation(ctx, req)
}

func DeleteEducation(ctx context.Context, req *system.SystemEducationRequest) (resp bool, err error) {
	return service.DeleteEducation(ctx, req)
}
