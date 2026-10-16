package handle

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryAreaPage(ctx context.Context, request *system.SystemAreaRequest) (*system.SystemAreaPageResponse, error) {
	return service.QueryAreaPage(ctx, request)
}

func QueryAreaTree(ctx context.Context, request *system.SystemAreaRequest) ([]*system.SystemAreaResponse, error) {
	return service.QueryAreaTree(ctx, request)
}

func QueryAreaList(ctx context.Context, request *system.SystemAreaRequest) ([]*system.SystemAreaResponse, error) {
	return service.QueryAreaList(ctx, request)
}

func QueryArea(ctx context.Context, request *system.SystemAreaRequest) (*system.SystemAreaResponse, error) {
	return service.QueryArea(ctx, request)
}

func SaveArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	return service.SaveArea(ctx, request)
}

func UpdateArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	return service.UpdateArea(ctx, request)
}

func DeleteArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	return service.DeleteArea(ctx, request)
}
