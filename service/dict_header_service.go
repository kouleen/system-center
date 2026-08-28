package service

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/repository"
)

func QueryDictHeaderPage(ctx context.Context, req *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderPageResponse, err error) {
	list, total, err := repository.QueryDictHeaderPage(ctx, req)
	systemDictHeaderList := make([]*system.SystemDictHeaderResponse, len(list))
	for index, dictHeader := range list {
		var updateTime int64
		if dictHeader.UpdateTime != nil {
			updateTime = dictHeader.UpdateTime.UnixMilli()
		}
		header := &system.SystemDictHeaderResponse{
			BaseResponse: &common.BaseResponse{
				Id:         dictHeader.ID,
				IsDelete:   int32(dictHeader.IsDelete),
				CreatedBy:  dictHeader.CreatedBy,
				UpdatedBy:  dictHeader.UpdatedBy,
				CreateTime: dictHeader.CreateTime.UnixMilli(),
				UpdateTime: updateTime,
			},
			DictName: dictHeader.DictName,
			DictType: dictHeader.DictType,
			Status:   common.BaseStatus(dictHeader.Status),
			Remark:   dictHeader.Remark,
		}
		systemDictHeaderList[index] = header
	}
	return &system.SystemDictHeaderPageResponse{
		Total:   total,
		Records: systemDictHeaderList,
	}, err
}
