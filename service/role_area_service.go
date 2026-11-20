package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/kouleen/common/pkg/ctxutil"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/modle"
	"github.com/kouleen/system-center/repository"
)

func QueryRoleAreaList(ctx context.Context, req *system.SystemRoleAreaRequest) (resp []*system.SystemRoleAreaResponse, err error) {
	list, err := repository.QueryRoleAreaByRoleId(ctx, req.GetRoleId())
	if err != nil {
		return
	}
	records := make([]*system.SystemRoleAreaResponse, len(list))
	for index, dictLine := range list {
		records[index] = dictLine.ConvertResp()
	}
	return records, nil
}

func SaveRoleArea(ctx context.Context, req *system.SystemRoleAreaRequest) (resp bool, err error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return false, err
	}
	areaIdList := req.GetAreaIds()
	systemRoleAreas := make([]*modle.SystemRoleArea, 0, len(areaIdList))
	if len(areaIdList) > 0 {
		for _, areaId := range areaIdList {
			systemRoleArea := &modle.SystemRoleArea{
				ID:        node.Generate().Int64(),
				AreaId:    areaId,
				RoleId:    req.GetRoleId(),
				CreatedBy: ctxutil.GetUserId(ctx),
				UpdatedBy: ctxutil.GetUserId(ctx),
			}
			systemRoleAreas = append(systemRoleAreas, systemRoleArea)
		}
	}
	if err = repository.SaveRoleArea(ctx, req.GetRoleId(), systemRoleAreas); err != nil {
		return
	}
	return true, nil
}
