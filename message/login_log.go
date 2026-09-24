package message

import (
	"context"
	"encoding/json"

	"github.com/kouleen/common/message"
	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
	"github.com/kouleen/system-center/utils"
)

type LoginLogConsumer struct{}

func (LoginLogConsumer) Queue() string {
	return "user.login.log.queue"
}

func (LoginLogConsumer) HandleMessage(ctx context.Context, msg *message.Message) error {
	var req system.SystemLoginLogRequest
	if err := json.Unmarshal(msg.Body, &req); err != nil {
		return err
	}
	location := utils.GetIPLocationOnline(ctx, req.Ip)
	req.Location = location
	if _, err := service.CreateLoginLog(ctx, &req); err != nil {
		return err
	}
	return nil
}

func init() {
	message.MustRegisterConsumer(new(LoginLogConsumer))
}
