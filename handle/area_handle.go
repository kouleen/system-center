package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryAreaPage(ctx context.Context, request *system.SystemAreaRequest) (*system.SystemAreaPageResponse, error) {
	return service.QueryAreaPage(ctx, request)
}

func QueryAreaList(ctx context.Context, request *system.SystemAreaRequest) ([]*system.SystemAreaResponse, error) {
	if request.GetParentCode() == "" && request.GetAreaCode() == "" && request.GetRegionCode() == "" && request.GetAreaLevelCode() == "" {
		return nil, errors.New("parent code can not be empty")
	}
	return service.QueryAreaList(ctx, request)
}

func QueryArea(ctx context.Context, request *system.SystemAreaRequest) (*system.SystemAreaResponse, error) {
	return service.QueryArea(ctx, request)
}

func SaveArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	if err := checkCreateArea(request); err != nil {
		return false, err
	}
	return service.SaveArea(ctx, request)
}

func UpdateArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	if err := checkUpdateArea(request); err != nil {
		return false, err
	}
	return service.UpdateArea(ctx, request)
}

func DeleteArea(ctx context.Context, request *system.SystemAreaRequest) (bool, error) {
	if len(request.GetIdList()) == 0 {
		return false, errors.New("id list is empty")
	}
	return service.DeleteArea(ctx, request)
}

func checkCreateArea(req *system.SystemAreaRequest) error {
	if req.GetAreaCode() == "" {
		return errors.New("area code can not be empty")
	}
	if req.GetAreaName() == "" {
		return errors.New("area name can not be empty")
	}
	if req.GetAreaDesc() == "" {
		return errors.New("area desc can not be empty")
	}
	if req.GetRegionCode() == "" {
		return errors.New("region code can not be empty")
	}
	if req.GetRegionName() == "" {
		return errors.New("region name can not be empty")
	}
	if req.GetRegionDesc() == "" {
		return errors.New("region desc can not be empty")
	}
	if req.GetCountryCode() == "" {
		return errors.New("country code can not be empty")
	}
	if req.GetCountryName() == "" {
		return errors.New("country name can not be empty")
	}
	if req.GetCountryDesc() == "" {
		return errors.New("country desc can not be empty")
	}
	if req.GetAreaLevelCode() == "" {
		return errors.New("area level code can not be empty")
	}
	return nil
}

func checkUpdateArea(req *system.SystemAreaRequest) error {
	if err := checkCreateArea(req); err != nil {
		return err
	}
	if req.Id == nil {
		return errors.New("id is required")
	}
	return nil
}
