package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryBulletinPage(ctx context.Context, req *system.SystemBulletinRequest) (resp *system.SystemBulletinPageResponse, err error) {
	return service.QueryBulletinPage(ctx, req)
}

func QueryBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp *system.SystemBulletinResponse, err error) {
	return service.QueryBulletin(ctx, req)
}

func SaveBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	if err = checkCreateBulletin(req); err != nil {
		return false, err
	}
	return service.SaveBulletin(ctx, req)
}

func UpdateBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	if err = checkUpdateBulletin(req); err != nil {
		return false, err
	}
	return service.UpdateBulletin(ctx, req)
}

func DeleteBulletin(ctx context.Context, req *system.SystemBulletinRequest) (resp bool, err error) {
	if req.GetIdList() == nil || len(req.GetIdList()) == 0 {
		return false, errors.New("id is required ")
	}
	return service.DeleteBulletin(ctx, req)
}

func checkCreateBulletin(req *system.SystemBulletinRequest) error {
	if req.GetTitle() == "" {
		return errors.New("title is required")
	}
	if req.Type == nil {
		return errors.New("type is required")
	}
	if req.Status == nil {
		return errors.New("status is required")
	}
	if req.GetContent() == "" {
		return errors.New("content is required")
	}
	return nil
}

func checkUpdateBulletin(req *system.SystemBulletinRequest) error {
	if req.Id == nil {
		return errors.New("id is required")
	}
	return checkCreateBulletin(req)
}
