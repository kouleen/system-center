package message

import (
	"context"
	"encoding/json"

	"github.com/kouleen/common/message"
	"github.com/kouleen/common/pkg/logutil"
	"github.com/kouleen/idl/kitex_gen/common"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

type DDLLog struct{}

func (DDLLog) Queue() string {
	return common.SYSTEM_CHANGE_LOG_QUEUE
}

func (DDLLog) HandleMessage(ctx context.Context, msg *message.Message) error {
	defer func() {
		if err := recover(); err != nil {
			logutil.CtxError(ctx, "recover panic err: %#v", err)
		}
	}()
	var req []*system.SystemChangeLogRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}

	if req[0].GetSchema() == "live_system" && req[0].GetTable() == "system_change_log" {
		return nil
	}
	if err := service.BatchCreateChangeLog(ctx, req); err != nil {
		return err
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(DDLLog))
}
