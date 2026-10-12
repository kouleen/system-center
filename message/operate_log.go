package message

import (
	"context"
	"encoding/json"

	"github.com/kouleen/common/message"
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
	var req system.SystemOperateLogRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	location := utils.GetIPLocationOnline(ctx, req.Ip)
	req.Location = location
	if _, err := service.CreateOperateLog(ctx, &req); err != nil {
		return err
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(OperateLogConsumer))
}
