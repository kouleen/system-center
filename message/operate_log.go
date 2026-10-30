package message

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/logutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
	"github.com/kouleen/system-center/utils"
)

type OperateLogConsumer struct{}

func (OperateLogConsumer) Queue() string {
	return common.SYSTEM_OPERATE_LOG_QUEUE
}

func (OperateLogConsumer) HandleMessage(ctx context.Context, msg *message.Message) error {
	defer func() {
		if err := recover(); err != nil {
			logutil.CtxError(ctx, "recover panic err: %#v", err)
		}
	}()
	var req system.SystemOperateLogRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	if err := checkOperateLogReq(&req); err != nil {
		logger.CtxErrorf(ctx, "checkOperateLogReq error: %v", err)
		return nil
	}
	location := utils.GetIPLocationOnline(ctx, req.Ip)
	req.Location = location
	if _, err := service.CreateOperateLog(ctx, &req); err != nil {
		return err
	}
	return nil
}

func checkOperateLogReq(req *system.SystemOperateLogRequest) error {
	if req.GetModule() == "" {
		return errors.New("module is empty")
	}
	if req.GetBusiness() == "" {
		return errors.New("business is empty")
	}
	if req.GetType() == "" {
		return errors.New("type is empty")
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(OperateLogConsumer))
}
